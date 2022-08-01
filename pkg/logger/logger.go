package logger

import (
	"encoding/json"
	"fmt"
	"go-devops-admin/pkg/app"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 全局Logger对象
var Logger *zap.Logger

// 初始化日志配置
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	// 获取日志介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// 设置日志等级
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别错误. 请修改config/log.go文件中的log.level配置项")
	}

	// 初始化core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化Logger
	Logger = zap.New(core,
		zap.AddCaller(),                   // 调用文件和行号, 内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层, 调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	// 将自定义的logger替换全局的logger
	// zap.L().Fatal()调用时，就会使用我们自定义的Logger
	zap.ReplaceGlobals(Logger)
}

// 设置日志存储格式
func getEncoder() zapcore.Encoder {
	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,        // 日志行尾追加 \n 换行符
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 日志级别名称大写 ERROR,INFO
		EncodeTime:     customTimerEncoder,               // 时间格式 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder,   // 执行时间, 以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,       // caller短格式， 长格式为绝对路径
	}

	// 本地环境配置
	if app.IsLocal() {
		// 终端输出关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func customTimerEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// 日志记录介质. GORM中使用两种介质，os.Stdout 和 文件
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	// 配置了按照日期记录日志文件
	if logType == "daily" {
		logName := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logName)
	}

	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		Compress:   compress,
	}

	// 配置输出介质
	if app.IsLocal() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 生产环境只记录日志文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

/* 添加一些日志辅助方法，既方便我们的调用，又对 zap 进行封装 */
// Dump 调试专用,不会中断程序,会在终端打印warning消息
// 第一个参数将内容转换成json格式，第二个参数可选,添加消息title
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

// 当 err != nil 时, 记录不同等级日志
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred", zap.Error(err))
	}
}

func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred", zap.Error(err))
	}
}

func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred", zap.Error(err))
	}
}

// DEBUG 调试日志,详尽的程序日志(不建议在生成环境使用)
// 调试示例:
//     logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

// INFO 通知类日志，一般是正常的日志
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// Warn类日志
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error类日志
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// 记录一条字符串类型的日志
//         logger.DebugString("SMS","短信内容",string(result.RowResponse))
func DebugString(moduleName string, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName string, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName string, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName string, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

// 记录一条 JSON 格式的日志
//         logger.DebugString("SMS","短信内容",string(result.RowResponse))
func DebugJson(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJson(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJson(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJson(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

// 格式化JSON字符串
func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON Marshal error", err.Error()))
	}
	return string(b)
}

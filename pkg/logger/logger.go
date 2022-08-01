package logger

import (
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
	Logger := zap.New(core,
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

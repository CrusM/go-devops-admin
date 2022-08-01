package logger

import (
	"context"
	"errors"
	"go-devops-admin/pkg/helpers"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// GormLogger 操作对象,实现gormlogger.interface
type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

// 外部调用. 实例化一个GormLogger对象. 示例：
// DB, err = gorm.Open(config, &gorm.Config{
// 		Logger: logger.NewGormLogger(),
// })
func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,                 // 使用全局logger.Logger对象
		SlowThreshold: 200 * time.Millisecond, // 慢查询阈值, 单位毫秒
	}
}

// LogMode 实现 gormlogger.Interface{} 的LogMode方法
func (l GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

// 实现gormlogger.Interface{} 的 Error 方法
func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Debugf(str, args...)
}

// 实现gormlogger.Interface{} 的 Warn 方法
func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Warnf(str, args...)
}

// 实现gormlogger.Interface{} 的 Info 方法
func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	l.logger().Sugar().Infof(str, args...)
}

// 实现gormlogger.Interface{} 的 Trace 方法
func (l GormLogger) Trace(ctx context.Context, begin time.Time, fn func() (string, int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 请求和返回条目
	sql, rows := fn()

	// 通用字段
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", helpers.MicrosecondsStr(elapsed)),
		zap.Int64("rows", rows),
	}

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误用 error 等级
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow Log", logFields...)
	}

	// 记录所有的 SQL 请求
	l.logger().Debug("Database Query", logFields...)
}

// logger内的辅助方法，确保Zap内置信息 Caller 的准确性(如：paginator/paginator.go:148)
func (l GormLogger) logger() *zap.Logger {
	// 跳过 gorm 内置调用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapGormPackage = filepath.Join("moul.io", "zapgorm2")
	)
	// 减去一次封装, 以及一次在 logger 初始化里添加的 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapGormPackage):
		default:
			// 返回一个附带跳过行号的新 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}

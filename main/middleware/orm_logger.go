package middleware

import (
	"context"
	"regexp"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

// ZapGormLogger 完全实现 logger.Interface 接口（保留原名）
type ZapGormLogger struct {
	ZapLogger     *zap.Logger
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// NewZapGormLogger 创建新的 ZapGormLogger 实例（保留原名）
func NewZapGormLogger(zapLogger *zap.Logger) *ZapGormLogger {
	return &ZapGormLogger{
		ZapLogger:     zapLogger,
		LogLevel:      logger.Info,
		SlowThreshold: 200 * time.Millisecond,
	}
}

// LogMode 设置日志级别（保留原名）
func (l *ZapGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 实现 logger.Interface 的 Info 方法（保留原名）
func (l *ZapGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.ZapLogger.Sugar().Infof(msg, data...)
	}
}

// Warn 实现 logger.Interface 的 Warn 方法（保留原名）
func (l *ZapGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.ZapLogger.Sugar().Warnf(msg, data...)
	}
}

// Error 实现 logger.Interface 的 Error 方法（保留原名）
func (l *ZapGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.ZapLogger.Sugar().Errorf(msg, data...)
	}
}

// Trace 实现 logger.Interface 的 Trace 方法
func (l *ZapGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()
	maskedSQL := maskSQL(sql)

	// 确保所有SQL查询都能被记录（无论是否错误）
	if l.LogLevel >= logger.Info {
		fields := []zap.Field{
			zap.Duration("elapsed", elapsed),
			zap.Int64("rows", rows),
			zap.String("sql", maskedSQL),
		}

		switch {
		case err != nil:
			l.ZapLogger.Error("ERROR", append(fields, zap.Error(err))...)
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
			l.ZapLogger.Warn("WARN", fields...)
		default:
			// 使用Info级别确保一定会输出
			l.ZapLogger.Info("INFO", fields...)
		}
	}
}

// ParamsFilter 实现 logger.Interface 的 ParamsFilter 方法（保留原名）
func (l *ZapGormLogger) ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{}) {
	// 如果需要过滤敏感参数，可以在这里实现
	return sql, params
}

// maskSQL 敏感信息过滤（新增辅助函数）
func maskSQL(sql string) string {
	return regexp.MustCompile(`(password|token|secret)=('[^']*'|"[^"]*")`).
		ReplaceAllString(sql, "$1=***")
}

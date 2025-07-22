package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 处理请求

		// 获取业务状态码（默认200）
		bizCode := http.StatusOK
		if code, exists := c.Get("code"); exists {
			bizCode = code.(int)
		}

		// 动态设置日志级别
		fields := []zap.Field{
			zap.Int("http_status", bizCode),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.ClientIP()),
			zap.Duration("latency", time.Since(start)),
		}

		switch {
		case bizCode >= 500:
			logger.Error("ERROR", fields...)
		case bizCode >= 400:
			logger.Warn("WARN", fields...)
		default:
			logger.Info("INFO", fields...)
		}

	}
}

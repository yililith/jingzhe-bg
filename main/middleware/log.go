package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"jingzhe-bg/main/internal/log"
	"jingzhe-bg/main/utils"
	"math"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	maxRequestBodySize    = 1 << 20  // 1MB
	maxResponseBodyToLog  = 4096     // 4KB
	multipartFormMemLimit = 32 << 20 // 32MB
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

var (
	respPool = sync.Pool{
		New: func() interface{} {
			return &utils.Response{}
		},
	}

	blwPool = sync.Pool{
		New: func() interface{} {
			return &bodyLogWriter{
				body: bytes.NewBuffer(make([]byte, 0, 1024)),
			}
		},
	}
)

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	if w.body == nil {
		w.body = bytes.NewBuffer(make([]byte, 0, 1024))
	}
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyLogWriter) reset() {
	w.ResponseWriter = nil
	if w.body != nil {
		w.body.Reset()
	}
}

func getClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func formatSize(size int64) string {
	if size == 0 {
		return "0"
	}
	if size == math.MinInt64 {
		return "-9223372036854775808"
	}

	var buf [20]byte
	i := len(buf)
	neg := size < 0
	if neg {
		size = -size
	}

	for size > 0 {
		i--
		buf[i] = byte('0' + size%10)
		size /= 10
	}

	if neg {
		i--
		buf[i] = '-'
	}

	return string(buf[i:])
}

func ZapLogger() gin.HandlerFunc {
	const (
		msgSuccess      = "request success"
		msgClientError  = "request client error"
		msgServerError  = "request server error"
		msgUnknown      = "request"
		fileUploadLabel = "file upload: "
		sizeLabel       = ", size: "
	)

	return func(c *gin.Context) {

		start := time.Now()
		clientIP := getClientIP(c.Request)
		var requestBody string

		// Request body processing
		if c.Request.ContentLength > 0 && c.Request.ContentLength < maxRequestBodySize {
			contentType := c.Request.Header.Get("Content-Type")
			switch {
			case strings.Contains(contentType, "multipart/form-data"):
				if err := c.Request.ParseMultipartForm(multipartFormMemLimit); err == nil {
					if file, header, err := c.Request.FormFile("file"); err == nil {
						defer file.Close()
						var b strings.Builder
						b.Grow(len(fileUploadLabel) + len(header.Filename) + len(sizeLabel) + 20)
						b.WriteString(fileUploadLabel)
						b.WriteString(header.Filename)
						b.WriteString(sizeLabel)
						b.WriteString(formatSize(header.Size))
						requestBody = b.String()
					}
				}
			default:
				bodyBytes, err := io.ReadAll(io.LimitReader(c.Request.Body, maxRequestBodySize))
				if err == nil {
					c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
					if len(bodyBytes) > 0 {
						if json.Valid(bodyBytes) {
							requestBody = string(bodyBytes)
						} else {
							requestBody = "[binary or non-json data]"
						}
					}
				}
			}
		}

		// Setup response writer
		blw := blwPool.Get().(*bodyLogWriter)
		blw.reset()
		blw.ResponseWriter = c.Writer
		c.Writer = blw
		defer func() {
			blw.reset()
			blwPool.Put(blw)
		}()

		// 安全执行业务逻辑
		c.Next()

		latency := time.Since(start)
		resp := respPool.Get().(*utils.Response)
		defer respPool.Put(resp)
		*resp = utils.Response{}

		var responseData interface{}
		statusCode := c.Writer.Status()
		bodyBytes := blw.body.Bytes()

		// Parse JSON response if applicable
		contentType := c.Writer.Header().Get("Content-Type")
		if strings.HasPrefix(contentType, "application/json") && len(bodyBytes) > 0 {
			if len(bodyBytes) <= maxResponseBodyToLog {
				if err := json.Unmarshal(bodyBytes, resp); err != nil {
					log.Logger.Debug("Failed to unmarshal response",
						zap.String("path", c.Request.URL.Path),
						zap.Error(err))
				} else {
					statusCode = resp.Code
					responseData = resp.Data
				}
			} else {
				responseData = "[response too large]"
			}
		}

		// Prepare log fields
		logFields := make([]zap.Field, 0, 12)
		logFields = append(logFields,
			zap.String("path", c.Request.URL.Path),
			zap.String("method", c.Request.Method),
			zap.String("ip", clientIP),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("latency", latency.String()),
			zap.String("timestamp", start.Format(time.RFC3339)),
		)

		if requestBody != "" {
			logFields = append(logFields, zap.String("request", requestBody))
		}

		if responseData != nil {
			logFields = append(logFields, zap.Reflect("response", responseData))
		}

		logFields = append(logFields,
			zap.Int("status", statusCode),
			zap.String("query", c.Request.URL.RawQuery),
		)

		// Log based on status code
		switch {
		case statusCode >= 200 && statusCode < 300:
			log.Logger.Info(msgSuccess, logFields...)
		case statusCode >= 400 && statusCode < 500:
			log.Logger.Warn(msgClientError, logFields...)
		case statusCode >= 500:
			// 500错误正常记录，但不触发panic
			log.Logger.Error(msgServerError, logFields...)
		default:
			log.Logger.Info(msgUnknown, logFields...)
		}
	}
}

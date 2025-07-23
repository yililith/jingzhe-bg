package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// Cors
//
//	@Description: 跨域中间件
//	@return gin.HandlerFunc
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,                                                // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},          // 允许的响应头
		AllowCredentials: true,                                                // 不允许包含凭据的请求
		MaxAge:           24 * time.Hour,
	})
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/utils/auth"
	"jingzhe-bg/main/utils/ead"
	"net/http"
	"strings"
)

// 统一的中断响应方法
func abortWithJSON(c *gin.Context, httpCode int, businessCode int, msg string) {
	c.JSON(httpCode, gin.H{
		"code": businessCode, // 业务状态码
		"msg":  msg,
	})
	c.Abort()
}

// AuthMiddleware
//
//	@Description: 认证中间件
//	@return gin.HandlerFunc
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			abortWithJSON(c, http.StatusUnauthorized, http.StatusUnauthorized, "Authorization header required")
			return
		}

		// 支持格式：Bearer <token>
		if !strings.HasPrefix(authHeader, "Bearer ") {
			abortWithJSON(c, http.StatusUnauthorized, http.StatusUnauthorized, "invalid token format")
			return
		}

		jwtToken := strings.TrimSpace(authHeader[len("Bearer "):])
		if jwtToken == "" {
			abortWithJSON(c, http.StatusUnauthorized, http.StatusUnauthorized, "token is empty")
			return
		}

		// 解密 token（注意：RSA 性能低，高并发场景建议替代或缓存）
		decryptedToken, err := ead.DecryptWithPrivateKey(global.GVA_PRIVATE_KEY, jwtToken)
		if err != nil {
			abortWithJSON(c, http.StatusInternalServerError, http.StatusInternalServerError, "token decryption failed: "+err.Error())
			return
		}

		// 解析 JWT
		claims, parseErr := auth.ParseToken(decryptedToken)
		if parseErr != nil {
			abortWithJSON(c, http.StatusUnauthorized, http.StatusUnauthorized, "invalid token: "+parseErr.Error())
			return
		}

		// 将 UID 注入上下文
		c.Set("uid", claims.UID)

		c.Next()
	}
}

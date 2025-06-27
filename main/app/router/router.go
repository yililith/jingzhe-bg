package router

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/internal/config"
	"jingzhe-bg/main/middleware"
	"net"
	"strconv"
)

var (
	routerAuthRole   = make([]func(*gin.RouterGroup), 0)
	routerNoAuthRole = make([]func(*gin.RouterGroup), 0)
)

func InitRouter() {
	// 设置为 Release 模式（禁用 Debug 日志）
	gin.SetMode(gin.ReleaseMode)
	// 初始化路由
	router := gin.New()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    404,
			"message": "path not found",
		})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    404,
			"message": "method not found",
		})
	})
	// 解决跨域问题
	router.Use(middleware.Cors())
	router.Use(middleware.ZapLogger())
	// 认证角色路由
	authGroup := router.Group("/api/v1")
	authGroup.Use(middleware.AuthMiddleware())
	for _, r := range routerAuthRole {
		r(authGroup)
	}
	// 非认证角色路由
	noAuthGroup := router.Group("/api/v2")
	for _, r := range routerNoAuthRole {
		r(noAuthGroup)
	}
	// 端口启动配置
	server := config.AppConfig.Server

	address := net.JoinHostPort(server.Host, strconv.Itoa(server.Port))
	// 启动服务
	router.Run(address)
}

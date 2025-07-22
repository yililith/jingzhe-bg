package router

import (
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/middleware"
	"net"
	"strconv"
)

var (
	routerAuthRole   = make([]func(*gin.RouterGroup), 0)
	routerNoAuthRole = make([]func(*gin.RouterGroup), 0)
)

func InitRouter() {
	// 设置为 Release 模式（禁用Gin自带 Debug 日志）
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
	router.Use(middleware.ZapLogger(global.GVA_LOGGER))

	// 认证角色路由
	routerAuthGroup(router)
	// 非认证角色路由
	routerNoAuthGroup(router)

	// 端口启动配置
	server := global.GVA_CONFIG.Server

	address := net.JoinHostPort(server.Host, strconv.Itoa(server.Port))
	// 启动服务
	if err := router.Run(address); err != nil {
		panic(err)
	}
}

// 需认证路由
func routerAuthGroup(router *gin.Engine) {
	authGroup := router.Group("/api/v1")
	authGroup.Use(middleware.AuthMiddleware())
	for _, r := range routerAuthRole {
		r(authGroup)
	}
}

// 不需要认证路由
func routerNoAuthGroup(router *gin.Engine) {
	noAuthGroup := router.Group("/api/v2")
	for _, r := range routerNoAuthRole {
		r(noAuthGroup)
	}
}

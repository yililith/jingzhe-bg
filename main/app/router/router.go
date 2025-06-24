package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jingzhe-bg/main/internal/config"
	"jingzhe-bg/main/middleware"
)

var (
	routerAuthRole   = make([]func(*gin.RouterGroup), 0)
	routerNoAuthRole = make([]func(*gin.RouterGroup), 0)
)

func InitRouter() {
	// 初始化路由
	router := gin.Default()
	// 解决跨域问题
	router.Use(middleware.Cors())
	// 认证角色路由
	authGroup := router.Group("/api/v1")
	for _, r := range routerAuthRole {
		r(authGroup)
	}
	// 非认证角色路由
	noAuthGroup := router.Group("/api/v2")
	for _, r := range routerNoAuthRole {
		r(noAuthGroup)
	}
	// 端口启动配置
	conf := config.AppConfig

	address := fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port)
	// 启动服务
	router.Run(address)
}

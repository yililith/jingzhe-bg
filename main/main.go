package main

import (
	"jingzhe-bg/main/app/router"
	"jingzhe-bg/main/internal"
)

func main() {
	// 初始化配置文件,数据库链接,其他服务链接等
	internal.InitRun()
	// Gin启动
	router.InitRouter()
}

//package main
//
//import (
//	"github.com/gin-contrib/zap"
//	"github.com/gin-gonic/gin"
//	"go.uber.org/zap"
//	"time"
//)
//
//func main() {
//	gin.SetMode(gin.ReleaseMode)
//	// 初始化 Zap 日志
//	logger, _ := zap.NewProduction()
//	defer logger.Sync() // 确保日志刷新到磁盘
//
//	// 创建 Gin 引擎（禁用默认日志）
//	router := gin.New()
//
//	// 添加 Zap 中间件（替换 Gin 默认日志）
//	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
//
//	// 可选：添加 Recovery 中间件（用 Zap 记录 Panic）
//	router.Use(ginzap.RecoveryWithZap(logger, true))
//
//	// 路由示例
//	router.GET("/ping", func(c *gin.Context) {
//		//panic("pong")
//		c.String(200, "pong")
//	})
//
//	router.Run(":8080")
//}

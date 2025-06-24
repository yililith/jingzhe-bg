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

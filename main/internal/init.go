package internal

import (
	"jingzhe-bg/main/app/router"
	"jingzhe-bg/main/internal/config"
	"jingzhe-bg/main/internal/database"
	"jingzhe-bg/main/internal/log"
	"jingzhe-bg/main/internal/minio"
	"jingzhe-bg/main/internal/rsa"
	"jingzhe-bg/main/internal/validator"
)

// 对配置文件,数据库,对象存储,日志等进行初始化
func InitRun() {
	if err := config.InitConfig(); err != nil {
		panic(err.Error())
	}
	// 日志初始化
	log.InitLogger()
	// 数据库链接初始化
	if err := database.InitDB(); err != nil {
		panic(err.Error())
	}
	// 对象存储初始化
	if err := minio.InitMinio(); err != nil {
		panic(err.Error())
	}
	// 加载加密密钥
	if err := rsa.InitKey(); err != nil {
		panic(err.Error())
	}
	// 字符串验证
	validator.InitValidator()
	// 启动Gin服务
	router.InitRouter()
}

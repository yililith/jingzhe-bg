package internal

import (
	"fmt"
	"jingzhe-bg/main/internal/config"
	"jingzhe-bg/main/internal/database"
	"jingzhe-bg/main/internal/log"
	"jingzhe-bg/main/internal/rsa"
)

// 对配置文件,数据库,对象存储,日志等进行初始化
func InitRun() {
	if config_err := config.InitConfig(); config_err != nil {
		fmt.Println(config_err.Error())
	}
	// 日志初始化
	log.InitLog()
	// 数据库链接初始化
	if database_err := database.InitDB(); database_err != nil {
		fmt.Println(database_err.Error())
	}
	// 加载加密密钥
	if rsa_err := rsa.InitKey(); rsa_err != nil {
		fmt.Println(rsa_err.Error())
	}
}

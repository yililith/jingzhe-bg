package internal

import (
	"fmt"
	"jingzhe-bg/main/internal/config"
	"jingzhe-bg/main/internal/log"
)

// 对配置文件,数据库,对象存储,日志等进行初始化
func InitRun() {
	if config_err := config.InitConfig(); config_err != nil {
		fmt.Println(config_err.Error())
	}
	log.InitLog()
}

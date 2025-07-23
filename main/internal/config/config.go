package config

import (
	"github.com/spf13/viper"
	"jingzhe-bg/main/global"
)

// InitConfig
//
//	@Description: 配置初始化
//	@return er
func InitConfig() error {

	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv() //读取环境变量

	// 将配置文件读取到全局变量里
	if err := viper.Unmarshal(&global.GVA_CONFIG); err != nil {
		return err
	}

	return nil
}

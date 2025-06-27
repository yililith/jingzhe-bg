package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type conf struct {
	Server   serviceConfig  `mapstructure:"server"`
	Database databaseConfig `mapstructure:"database"`
	// Minio    minioConfig    `mapstructure:"minio"`
}

type serviceConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type databaseConfig struct {
	DBName   string `mapstructure:"db_name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type minioConfig struct {
	EndPoInt        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	BucketName      string `mapstructure:"bucketName"`
	FilePath        string `mapstructure:"filePath"`
	UseSLL          bool   `mapstructure:"useSLL"`
}

var AppConfig *conf

func InitConfig() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("[kc_blog_bkg_sys] [ERROR]: 配置文件读取失败")
	}

	viper.AutomaticEnv() //读取环境变量

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("[kc_blog_bkg_sys] [ERROR]: 配置文件映射失败")
	}
	return nil
}

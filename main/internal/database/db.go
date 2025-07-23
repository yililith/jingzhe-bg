package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jingzhe-bg/main/global"
	"jingzhe-bg/main/middleware"
	"net/url"
	"time"
)

var (
	maxRetries = 3
	retryDelay = 5 * time.Second
)

// InitDB
//
//	@Description: 数据库初始化
//	@return er
func InitDB() error {

	conf := global.GVA_CONFIG.Database

	// 使用更安全的DSN构建方式，避免密码中的特殊字符问题
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		url.QueryEscape(conf.Password), // 对密码进行URL编码
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	var err error
	var db *gorm.DB

	gormLogger := middleware.NewZapGormLogger(global.GVA_LOGGER)
	// 添加重试逻辑
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger,
			// 添加更多配置
			PrepareStmt:            true, // 预编译SQL
			SkipDefaultTransaction: true, // 跳过默认事务
		})

		if err == nil {
			break
		}

		if i < maxRetries-1 {
			time.Sleep(retryDelay)
		}
	}

	if err != nil {
		return err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	// 将连接池赋值给全局变量
	global.GVA_DB = db

	return nil
}

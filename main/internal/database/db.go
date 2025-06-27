package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jingzhe-bg/main/internal/config"
	"jingzhe-bg/main/internal/log"
	"jingzhe-bg/main/middleware"
	"net/url"
	"time"
)

var (
	DB         *gorm.DB
	maxRetries = 3
	retryDelay = 5 * time.Second
)

func InitDB() error {
	conf := config.AppConfig.Database

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

	gormLogger := middleware.NewZapGormLogger(log.Logger)
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
		return fmt.Errorf("failed to connect database after %d attempts: %v", maxRetries, err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	DB = db
	return nil
}

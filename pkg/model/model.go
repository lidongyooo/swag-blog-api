package model

import (
	"fmt"
	"github.com/lidongyooo/swag-blog-api/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB  {
	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = config.Viper.GetString("DB_HOST")
		port     = config.Viper.GetString("DB_PORT")
		database = config.Viper.GetString("DB_DATABASE")
		username = config.Viper.GetString("DB_USERNAME")
		password = config.Viper.GetString("DB_PASSWORD")
		charset  = config.Viper.GetString("DB_CHARSET")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.Open(dsn)

	var level logger.LogLevel

	if config.Viper.GetBool("APP_DEBUG") {
		// 读取不到数据也会显示
		level = logger.Info
	} else {
		// 只有错误才会显示
		level = logger.Error
	}

	fmt.Println(level)
	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})

	if err != nil {
		panic(err.Error())
	}

	return DB
}

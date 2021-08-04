package bootstrap

import (
	"fmt"
	"github.com/lidongyooo/swag-blog-api/app/models/article"
	"github.com/lidongyooo/swag-blog-api/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// DB gorm.DB 对象
var db *gorm.DB

func SetupDB() {
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
		level = logger.Warn
	} else {
		// 只有错误才会显示
		level = logger.Error
	}

	// 准备数据库连接池
	db, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})

	if err != nil {
		panic(err.Error())
	}

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.Viper.GetInt("DB_MAX_OPEN_CONNECTIONS"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.Viper.GetInt("DB_MAX_IDLE_CONNECTIONS"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Viper.GetInt("DB_MAX_LIFE_SECONDS")) * time.Second)

	migration()
}

func migration() {
	db.AutoMigrate(
		&article.Article{},
		&article.Tag{},
	)
}

package bootstrap

import (
	"github.com/lidongyooo/swag-blog-api/app/models/article"
	"github.com/lidongyooo/swag-blog-api/pkg/config"
	"github.com/lidongyooo/swag-blog-api/pkg/model"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.Viper.GetInt("DB_MAX_OPEN_CONNECTIONS"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.Viper.GetInt("DB_MAX_IDLE_CONNECTIONS"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Viper.GetInt("DB_MAX_LIFE_SECONDS")) * time.Second)

	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&article.Article{},
		&article.Tag{},
	)
}

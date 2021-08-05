package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/bootstrap"
	"github.com/lidongyooo/swag-blog-api/pkg/config"
	"net/http"
	"time"
)

func main() {

	if !config.Viper.GetBool("APP_DEBUG") {
		gin.SetMode(gin.ReleaseMode)
	}

	bootstrap.SetupDB()
	r := bootstrap.SetupRoute()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Viper.GetInt("SERVER_HTTP_PORT")),
		Handler:        r,
		ReadTimeout:    time.Duration(config.Viper.GetInt64("SERVER_READ_TIMEOUT")) * time.Second,
		WriteTimeout:   time.Duration(config.Viper.GetInt64("SERVER_WRITE_TIMEOUT")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
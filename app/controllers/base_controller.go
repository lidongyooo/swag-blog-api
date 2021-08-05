package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/pkg/config"
	res "github.com/lidongyooo/swag-blog-api/pkg/response"
	"net/http"
)

type BaseController struct {
}

func (bc BaseController) ResponseForError(context *gin.Context, err error)  {
	if config.Viper.GetBool("APP_DEBUG") {
		context.JSON(http.StatusInternalServerError, res.New(res.ERROR, err.Error()))
	}

	context.JSON(http.StatusInternalServerError, res.New(res.ERROR, res.GetMsg(res.ERROR)))
}
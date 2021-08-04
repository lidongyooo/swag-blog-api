package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/app/models/article"
	res "github.com/lidongyooo/swag-blog-api/pkg/response"
	"net/http"
)

type ArticlesController struct {
}

type ArticleExtension struct {
	article.Article

	TagName string `json:"tag_name" binding:"required"`
}

func (articles *ArticlesController) Store (context *gin.Context)  {
	artExt := &ArticleExtension{}
	err := context.ShouldBindJSON(&artExt)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, res.Error(res.INVALID_PARAMS, err))
	} else {
		context.JSON(http.StatusCreated, artExt)
	}
}
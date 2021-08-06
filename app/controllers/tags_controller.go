package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/app/models/article"
	res "github.com/lidongyooo/swag-blog-api/pkg/response"
	"net/http"
)

type TagsController struct {
	BaseController
}

func (tagsCon *TagsController) Index(context *gin.Context) {
	tags, err := article.GetTagsAll()

	if err != nil {
		tagsCon.ResponseForError(context, err)
	} else {
		context.JSON(http.StatusOK, res.New(res.SUCCESS, res.GetMsg(res.SUCCESS)).WithData(tags))
	}

}
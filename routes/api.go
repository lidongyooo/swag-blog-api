package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/app/controllers"
)

func RegisterApiRoutes(r *gin.Engine)  {
	ac := new(controllers.ArticlesController)
	articles := r.Group("/articles")
	articles.GET("", ac.Index)
	articles.POST("", ac.Store)
	articles.POST("/:id", ac.Update)
	articles.GET("/:id", ac.Show)

	tc := new(controllers.TagsController)
	tags := r.Group("/tags")
	tags.GET("", tc.Index)
}

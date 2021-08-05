package routes

import (
	"github.com/gin-gonic/gin"
	controllers2 "github.com/lidongyooo/swag-blog-api/app/controllers"
)

func RegisterApiRoutes(r *gin.Engine)  {
	ac := new(controllers2.ArticlesController)
	articles := r.Group("/articles")
	articles.POST("", ac.Store)
	articles.POST("/:id", ac.Update)

}

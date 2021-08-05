package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/app/models/controllers"
)

func RegisterApiRoutes(r *gin.Engine)  {
	ac := new(controllers.ArticlesController)
	articles := r.Group("/articles")
	articles.POST("", ac.Store)
	articles.POST("/:id", ac.Update)

}

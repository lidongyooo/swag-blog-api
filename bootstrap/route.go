package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/lidongyooo/swag-blog-api/routes"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	routes.RegisterApiRoutes(r)
	return r
}

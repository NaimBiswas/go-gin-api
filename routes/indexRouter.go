package routes

import (
	"NaimBiswas/go-gin-api/routes/exportServices"
	UserRoutes "NaimBiswas/go-gin-api/routes/userRoutes"

	"github.com/gin-gonic/gin"
)

func MainRoutes(api *gin.RouterGroup) {
	UserRoutes.Main(api.Group("/user"))
	exportServices.Main(api.Group("/export"))
}

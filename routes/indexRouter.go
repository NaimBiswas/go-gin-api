package routes

import (
	"NaimBiswas/go-gin-api/routes/exportRouter"
	UserRoutes "NaimBiswas/go-gin-api/routes/userRoutes"

	"github.com/gin-gonic/gin"
)

func MainRoutes(api *gin.RouterGroup) {
	UserRoutes.Main(api.Group("/user"))
	exportRouter.Main(api.Group("/export"))
}

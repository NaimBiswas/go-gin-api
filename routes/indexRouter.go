package routes

import (
	UserRoutes "NaimBiswas/go-gin-api/routes/userRoutes"

	"github.com/gin-gonic/gin"
)

func MainRoutes(api *gin.RouterGroup) {
	UserRoutes.Main(api.Group("/user")) 
}

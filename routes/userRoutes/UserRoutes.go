package UserRoutes

import (
	"NaimBiswas/go-gin-api/controllers/userController"
	authMiddlewares "NaimBiswas/go-gin-api/middlewares"

	"github.com/gin-gonic/gin"
)

func Main(api *gin.RouterGroup) {
	api.Use(authMiddlewares.VerifyUser)
	api.GET("/", userController.GetUser)
	api.GET("/:id", userController.GetAUser)
	api.POST("/create", userController.CreateUser)
	api.PUT("/update/:id", userController.UpdateUser)
	api.DELETE("/delete/:id", userController.DeleteUser)
}

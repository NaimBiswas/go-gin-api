package UserRoutes

import (
	"NaimBiswas/go-gin-api/controllers/UserController"
	authMiddlewares "NaimBiswas/go-gin-api/middlewares"

	"github.com/gin-gonic/gin"
)

func Main(api *gin.RouterGroup) {
	api.Use(authMiddlewares.VerifyUser)
	api.GET("/", UserController.GetUser)
	api.GET("/:id", UserController.GetAUser)
	api.POST("/create", UserController.CreateUser)
	api.PUT("/update/:id", UserController.UpdateUser)
	api.DELETE("/delete/:id", UserController.DeleteUser)
}

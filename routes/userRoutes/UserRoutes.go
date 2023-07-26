package UserRoutes

import (
	"NaimBiswas/go-gin-api/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func Main(api *gin.RouterGroup) {
	api.GET("/", UserController.GetUser)
	api.GET("/:id", UserController.GetAUser)
	api.POST("/create", UserController.CreateUser)
	api.PUT("/update/:id", UserController.UpdateUser)
	api.DELETE("/delete/:id", UserController.DeleteUser)
}

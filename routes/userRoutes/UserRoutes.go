package UserRoutes

import (
	"NaimBiswas/go-gin-api/controllers/UserController"

	"github.com/gin-gonic/gin"
)

func Main(api *gin.RouterGroup) {
	api.GET("/get", UserController.GetUser)
	api.POST("/create", UserController.CreateUser)
	api.PUT("/update/:id", UserController.UpdateUser)
	api.DELETE("/delete/:id", UserController.DeleteUser)
}




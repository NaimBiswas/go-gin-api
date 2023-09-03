package productRouter

import (
	"NaimBiswas/go-gin-api/controllers/productController"
	"github.com/gin-gonic/gin"
)

func New(api *gin.RouterGroup) {
	api.GET("/", productController.GetAllProducts)
	api.GET("/imports/", productController.ImportedProducts)
}

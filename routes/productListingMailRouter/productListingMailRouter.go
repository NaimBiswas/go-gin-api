package productListingMailRouter

import (
	"NaimBiswas/go-gin-api/controllers/productMailController"
	"github.com/gin-gonic/gin"
)

func New(api *gin.RouterGroup) {
	api.POST("/fire-product-mail", productMailController.SendMail)
}

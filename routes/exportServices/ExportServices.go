package exportServices

import (
	"NaimBiswas/go-gin-api/controllers/exportController"
	"github.com/gin-gonic/gin"
)

func Main(api *gin.RouterGroup) {
	api.GET("/pdf/:name", exportController.ExportToPdf)
}

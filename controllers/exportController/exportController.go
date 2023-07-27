package exportController

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	"NaimBiswas/go-gin-api/services/exportServices"
	"github.com/gin-gonic/gin"
)

func ExportToPdf(c *gin.Context) {
	name := c.Param("name")
	collectionName := dbConfig.GetCollection(dbConfig.DB, name)

	exportServices.ExportToPdf(c, collectionName)
}

func ExportToXlsx(c *gin.Context) {

}

func ExportToCSV(c *gin.Context) {
	name := c.Param("name")
	collectionName := dbConfig.GetCollection(dbConfig.DB, name)

	exportServices.ExportToCSV(c, collectionName)
}

func CSVExportCore() {

}

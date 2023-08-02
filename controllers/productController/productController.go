package productController

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	response "NaimBiswas/go-gin-api/response"
	"NaimBiswas/go-gin-api/services/productServices"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var productCollection *mongo.Collection = dbConfig.GetCollection(dbConfig.DB, "products")

func GetAllProducts(c *gin.Context) {

	limit := c.Query("limit")
	page := c.Query("page")

	if limit == "" || page == "" {
		limit = "10"
		page = "1"
	}
	limitInNumber, err := strconv.Atoi(limit)
	pageInNumber, err := strconv.Atoi(page)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Fail to convert limit && page", err.Error())
		return
	}
	data, dataCount, err := productServices.GetAllProduct(limitInNumber, pageInNumber, productCollection)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Something went wrong in ProductService", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"results":      data,
		"success":      true,
		"limit":        limitInNumber,
		"page":         pageInNumber,
		"totalPages":   int(math.Ceil(float64(int(dataCount)) / float64(limitInNumber))),
		"totalResults": dataCount,
	})
}

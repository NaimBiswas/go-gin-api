package main

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	"NaimBiswas/go-gin-api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello Gin World!")

	Router := gin.Default()
	
    api := Router.Group("/api")
	routes.MainRoutes(api)

	Router.GET("/", func(c *gin.Context){
		c.JSON(200,"Welcome to gin world")
	})
	
	dbConfig.DbConnection()
	Router.Run(":3001")
}
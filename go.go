package main

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	"NaimBiswas/go-gin-api/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func main() {
	fmt.Println("Hello Gin World!")
	app = gin.New()
	Router := gin.Default()

	api := Router.Group("/api")
	routes.MainRoutes(api)

	Router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to gin world")
	})

	dbConfig.DbConnection()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	Router.Run(":" + port)
}

func handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

package main

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	"NaimBiswas/go-gin-api/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()
	Router.Use(cors.New(CORSConfig()))
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
func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

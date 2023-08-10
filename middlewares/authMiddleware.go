package authMiddlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Your are not authorize to access this resource",
			"message": "Something went wrong",
		})
		c.Abort()
		return
	} else {
		fmt.Println("toke:", token)
		c.Next()
	}
}

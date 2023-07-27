package authMiddlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyUser(c *gin.Context) {
	token := c.GetHeader("X-Authorization")
	if token == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Something went wrong",
			"message": "Your are not authorize to access this resource",
		})
		c.Abort()
		return
	} else {
		fmt.Println("toke:", token)
		c.Next()
	}
}

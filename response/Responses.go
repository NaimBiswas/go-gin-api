package Responses

import "github.com/gin-gonic/gin"

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func ErrorResponse(c *gin.Context, Status int, message string, errorMessage string) {
	c.JSON(Status, gin.H{
		"message": message, "error": errorMessage,
	})
	c.Abort()
}

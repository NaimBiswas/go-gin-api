package UserController

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	Models "NaimBiswas/go-gin-api/models"
	Responses "NaimBiswas/go-gin-api/response"
	CommonServices "NaimBiswas/go-gin-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = dbConfig.GetCollection(dbConfig.DB, "users")

func GetUser(c *gin.Context) {
	// Handle the GET request for "/user/get"

	var users []Models.User
	results, err := CommonServices.GetValues(userCollection, 10, 1)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "error", "Data": map[string]interface{}{"data": err.Error()}})
		return
	}
	defer results.Close(c)
	if err := results.All(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "error", "Data": map[string]interface{}{"data": err.Error()}})
		return
	}
	c.JSON(http.StatusOK,
		Responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"users": users}},
	)
}

func GetAUser(c *gin.Context) {
	UserId := c.Param("id")
	var user Models.User
	response := CommonServices.GetValueById(userCollection, 10, 1, UserId)

	if err := response.Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "error", "data": map[string]interface{}{"error": err.Error()}})
		return
	}
	if len(user.FirstName) > 0 {
		c.JSON(http.StatusOK, Responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"user": user}})
		return
	}
	c.JSON(http.StatusBadRequest, Responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"errMessage": "No data Found"}})

}

func CreateUser(c *gin.Context) {
	// Handle the POST request for "/user/create"
	c.JSON(200, gin.H{
		"message": "Create user",
	})
}

func UpdateUser(c *gin.Context) {
	// Handle the PUT request for "/user/update/:id"
	userID := c.Param("id")

	c.JSON(200, gin.H{
		"message": "Update user with ID: " + userID,
	})
}

func DeleteUser(c *gin.Context) {
	// Handle the DELETE request for "/user/delete/:id"
	userID := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Delete user with ID: " + userID,
	})
}

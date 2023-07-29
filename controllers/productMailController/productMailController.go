package productMailController

import (
	dbConfig "NaimBiswas/go-gin-api/DbConfig"
	CommonServices "NaimBiswas/go-gin-api/services"
	mailServices "NaimBiswas/go-gin-api/services/MailServices"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context) {
	//verify body obj and get the Object

	//get the user email list
	var userCollection *mongo.Collection = dbConfig.GetCollection(dbConfig.DB, "users")
	allMailAddress := CommonServices.GetUserMails(c, userCollection)
	fmt.Println("allMail:", allMailAddress)
	templatePath, _ := filepath.Abs("./htmlTemplates/mailTemplates/productCreatedMail.html")
	//send Mail
	res, err := mailServices.New("Product Title Will be there", allMailAddress, templatePath)
	if err != nil {
		c.JSON(500, gin.H{"message": "Something went wrong", "error": err.Error(), "success": false})
		c.Abort()
	}
	c.JSON(http.StatusAccepted, res)

}

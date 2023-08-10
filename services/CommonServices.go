package CommonServices

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Email string `bson:"email" json:"email"`
}

func GetValues(collection *mongo.Collection, pageSize int, pageNumber int) (*mongo.Cursor, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pagination := getPagination(pageSize, pageNumber)
	results, err := collection.Find(ctx, bson.M{}, pagination)
	dataCount, err := collection.CountDocuments(ctx, bson.M{})
	// var results *mongo.Cursor
	// var dataCount int64
	// var resultsErr error
	// var dataCountErr error

	// // Use a WaitGroup to wait for both goroutines to complete
	// var wg sync.WaitGroup
	// wg.Add(2)

	// // Fetch results using a goroutine
	// go func() {
	// 	defer wg.Done()
	// 	results, resultsErr = collection.Find(ctx, bson.M{}, pagination)
	// }()

	// // Fetch data count using a goroutine
	// go func() {
	// 	defer wg.Done()
	// 	dataCount, dataCountErr = collection.CountDocuments(ctx, bson.M{})
	// }()

	// // Wait for both goroutines to complete
	// wg.Wait()

	// // Check for errors
	// if resultsErr != nil {
	// 	return nil, 0, resultsErr
	// }
	// if dataCountErr != nil {
	// 	return nil, 0, dataCountErr
	// }


	return results, dataCount, err
}

func GetValueById(collection *mongo.Collection, pageSize, pageNumber int, id string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//convert Id to ObjectId
	objId, _ := primitive.ObjectIDFromHex(id)
	//pagination := getPagination(pageSize, pageNumber)
	results := collection.FindOne(ctx, bson.M{"_id": objId})
	return results
}

func getPagination(pageSize int, pageNumber int) *options.FindOptions {

	skip := (pageNumber - 1) * pageSize
	// Define the options for the query
	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))
	return findOptions
}

func GetUserMails(c *gin.Context, collection *mongo.Collection) []string {

	// Set up options to include the projection
	projection := bson.M{"email": 1}
	findOptions := options.Find().SetProjection(projection)

	result, err := collection.Find(c, bson.M{}, findOptions)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong to find the data", "error": err.Error()})
	}
	//defer result.Close(c)
	var emails []string

	// Iterate through the cursor and extract the email addresses
	for result.Next(c) {
		var user User
		err := result.Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong", "error": err.Error()})
		}
		emails = append(emails, user.Email)
	}
	return emails
}

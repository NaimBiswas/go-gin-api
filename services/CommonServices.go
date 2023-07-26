package CommonServices

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetValues(collection *mongo.Collection, pageSize int, pageNumber int) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pagination := getPagination(pageSize, pageNumber)
	results, err := collection.Find(ctx, bson.M{}, pagination)

	return results, err
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

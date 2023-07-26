package CommonServices

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetValues(collection *mongo.Collection, pageSize int, pageNumber int) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	skip := (pageNumber - 1) * pageSize

	// Define the options for the query
	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64(skip))
	results, err := collection.Find(ctx, bson.M{}, findOptions)

	return results, err
}

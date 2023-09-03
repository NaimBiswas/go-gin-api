package dbConfig

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetEnvValues(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	URL := os.Getenv(key)
	fmt.Println("URL:==========",URL)

	return URL
}

func DbConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetEnvValues("DATABASE_URL")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB================================================================")
	return client
}

// DB Client instance
var DB *mongo.Client = DbConnection()

// GetCollection getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := DB.Database(GetEnvValues("DATABASE_NAME")).Collection(collectionName)
	return collection
}

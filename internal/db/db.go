// internal/db/db.go

package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client

// Connect initializes a connection to the MongoDB database
func Connect() error {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(viper.GetString("mongodb.connection_string")).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		return err
	}

	client = c
	return nil
}

// GetClient returns the underlying mongo.Client object
func GetClient() *mongo.Client {
	return client
}

// GetCollection returns a mongo.Collection object for the specified collection
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

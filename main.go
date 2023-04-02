package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TheDevExperiment/server/router"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func startMongo() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(viper.GetString("mongodb.connection_string")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// TODO: use mongodb client returned below
	_, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

}

func setupViper() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func main() {
	setupViper() // FYI- This will always stay on top.
	startMongo()
	r := router.SetupRouter()
	r.Run(":8080")
}

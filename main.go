package main

import (
	"log"

	"github.com/TheDevExperiment/server/internal/cache"
	"github.com/TheDevExperiment/server/internal/db"
	logger "github.com/TheDevExperiment/server/internal/log"
	"github.com/TheDevExperiment/server/router"
	"github.com/spf13/viper"
)

func startMongo() {
	err := db.Connect()
	if err != nil {
		logger.Fatal(err)
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
		log.Fatalf("Error reading config file, %s", err)
	}
}

func startLogger() {
	// TODO: choose logger acc to env ie prod/dev/stag
	// default is dev
	logger.InitZapLogger()
}

func main() {
	setupViper()      // FYI- Be careful while adding anything above this.
	startLogger()     // start logger
	startMongo()      // start mongo db
	cache.GetClient() // fire up redis cache

	r := router.SetupRouter()
	r.Run(":8080")
}

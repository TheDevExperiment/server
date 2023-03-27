package main

import "github.com/TheDevExperiment/server/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}

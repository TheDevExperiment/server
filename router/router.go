package router

import (
	"github.com/TheDevExperiment/server/internal/auth"
	"github.com/TheDevExperiment/server/internal/db/repositories"
	"github.com/TheDevExperiment/server/middlewares"
	"github.com/gin-gonic/gin"
)

/*
	Pattern to follow here:
	- add all the routes in this file
	- create the respective models for Request and Response in the router/models package
	- Logic for reach route will be in a separate package. This package will refer to router/models for its usecase.

	I have created a simple auth route keeping this design in mind.
*/

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userRepository := repositories.NewUserRepository()

	// TODO: register middleware
	r.Use(func(c *gin.Context) {
		c.Set("userRepository", userRepository)
		c.Next()
	})

	// define all the routes
	r.POST("/auth/v1/guest-validate", middlewares.JWTAuthMiddleware, auth.GuestValidateV1)
	r.POST("/auth/v1/guest", auth.CreateGuestV1)
	return r
}

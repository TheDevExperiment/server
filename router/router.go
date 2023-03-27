package router

import (
	"net/http"

	"github.com/TheDevExperiment/server/internal/auth"
	"github.com/TheDevExperiment/server/router/models"
	"github.com/gin-gonic/gin"
)

/*
	Pattern to follow here:
	- add all the routes in this file
	- create the respective models for Request and Response in the router/models package
	- Logic for reach route will be in a separate package. This package will refer to router/models for its usecase.

	I have created a simple auth route keeping this design in mind.
*/

func guestValidateV1(c *gin.Context) {
	// first bind the req to our model
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//TODO: implement context cancellation(case where user cancels the request)
	res, err := auth.Handler(c, &req)

	if err == nil {
		c.JSON(http.StatusOK, res)
		return
	}

	// some error occurent
	c.String(http.StatusInternalServerError, "something went wrong")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// TODO: register middleware

	// define all the routes
	r.POST("/auth/v1/guestValidate", guestValidateV1)
	return r
}

package auth

import (
	"log"
	"net/http"

	"github.com/TheDevExperiment/server/internal/db/repositories"
	"github.com/TheDevExperiment/server/middlewares"
	"github.com/TheDevExperiment/server/router/models/auth"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GuestValidateV1(c *gin.Context) {
	var res auth.AuthResponse
	userRepository, ok := c.MustGet("userRepository").(*repositories.UserRepository)
	if !ok {
		res.Message = http.StatusText(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	//set by the middleware
	userId := c.GetString(middlewares.ContextKey_UserId)
	if userId == "" {
		res.Message = http.StatusText(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	data, err := userRepository.Find(c, bson.M{"_id": userId})
	if err != nil {
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if data == nil || len(data) < 1 {
		res.Message = http.StatusText(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	res.Message = "Authentication Successful"
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func CreateGuestV1(c *gin.Context) {
	// first bind the req to our model
	var req auth.CreateAccountRequest
	var res auth.AuthResponse
	userRepository, ok := c.MustGet("userRepository").(*repositories.UserRepository)
	if !ok {
		res.Message = http.StatusText(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Age == "" {
		req.Age = "18-20"
	}
	if req.CountryId == "" {
		req.CountryId = "IND"
	}
	if req.CityId == "" {
		req.CityId = "DEL"
	}
	data, err := userRepository.Create(c, req.Age, req.CountryId, req.CityId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Print(data)
	// some error occurent
	res.Message = "Created User."
	res.Data = []repositories.UserModel{data}
	c.JSON(http.StatusCreated, res)
}

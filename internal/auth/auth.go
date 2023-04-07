package auth

import (
	"log"
	"net/http"

	"github.com/TheDevExperiment/server/internal/db/repositories"
	"github.com/TheDevExperiment/server/router/models/authModel"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GuestValidateV1(c *gin.Context) {
	// first bind the req to our model
	var req authModel.AuthRequest
	var res authModel.AuthResponse
	userRepository, ok := c.MustGet("userRepository").(*repositories.UserRepository)
	if !ok {
		log.Fatal("oops!", ok)
		c.JSON(http.StatusInternalServerError, gin.H{"error": ok})
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if is not a guest account
	if !req.IsGuest {
		res.Message = "Only Guest Accounts are supported"
		c.JSON(http.StatusNotImplemented, res)
		return
	}
	if req.UserId == "" || req.SecretToken == "" {
		res.Message = "UserId and SecretToken must be provided."
		c.JSON(http.StatusBadRequest, res)
		return
	}

	hexId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := userRepository.Find(c, bson.M{"_id": hexId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Print(data)
	if data == nil || len(data) < 1 {
		res.Message = "Failed Authentication."
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	// if auth secret provided matches with the stored one,
	if req.SecretToken == data[0].GuestAuthSecret && req.UserId == data[0].Id.Hex() {
		res.Message = "User Verified Successfully."
		res.Data = data[0]
		c.JSON(http.StatusOK, res)
		return
	}

	// some error occurent
	res.Message = "Failed Authentication."
	c.JSON(http.StatusUnauthorized, res)
}

func CreateGuestV1(c *gin.Context) {
	// first bind the req to our model
	var req authModel.AuthRequest
	var res authModel.AuthResponse
	userRepository, ok := c.MustGet("userRepository").(*repositories.UserRepository)
	if !ok {
		log.Fatal("oops!", ok)
		c.JSON(http.StatusInternalServerError, gin.H{"error": ok})
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := userRepository.Create(c, "18-20", "IND", "DEL")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Print(data)
	// some error occurent
	res.Message = "Created User."
	res.Data = data
	c.JSON(http.StatusCreated, res)
}

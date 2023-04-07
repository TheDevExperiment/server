package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheDevExperiment/server/internal/db/repositories"
	"github.com/TheDevExperiment/server/router/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GuestValidateV1(c *gin.Context) {
	// first bind the req to our model
	var req models.AuthRequest
	var res models.AuthResponse
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

	data, err := userRepository.Find(c, bson.M{})
	fmt.Print(err)
	// some error occurent
	res.Message = "All Good, check server console for output."
	res.Data = data
	c.JSON(http.StatusOK, res)
}

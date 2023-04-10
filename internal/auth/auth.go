package auth

import (
	"fmt"
	"github.com/TheDevExperiment/server/internal/db/repositories"
	"github.com/TheDevExperiment/server/router/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GuestValidateV1(c *gin.Context) {
	// first bind the req to our model
	var req models.AuthRequest
	var res models.AuthResponse
	userRepository := c.MustGet("userRepository").(*repositories.UserRepository)

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

	allUsers, err := userRepository.Find(c, bson.M{})
	docs := make([]bson.D, len(allUsers))
	for i, r := range allUsers {
		docs[i] = r.(bson.D)
	}
	fmt.Println(docs, err)
	// some error occurent
	res.Message = "All Good, check server console for output."
	c.JSON(http.StatusOK, res)
}

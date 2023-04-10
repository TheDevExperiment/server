package post

import (
	"net/http"

	"github.com/TheDevExperiment/server/internal/db/posts"
	"github.com/TheDevExperiment/server/router/models"
	"github.com/gin-gonic/gin"
)

func CreateV1(c *gin.Context) {
	// first bind the req to our model
	var req models.CreatePostRequest
	var res models.CreatePostResponse
	pr := posts.NewPostsRepository()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := pr.Create(c, req.Post)
	if err != nil {
		httpCode := http.StatusInternalServerError
		res.Message = err.Error()
		c.JSON(httpCode, res)
		return
	}
	res.Message = "Created new post"
	c.JSON(http.StatusOK, res)
}

func NearbyV1(c *gin.Context) {
	// first bind the req to our model
	var req models.NearbyPostRequest
	var res models.NearbyPostResponse
	pr := posts.NewPostsRepository()

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts, err := pr.FindNearby(c, req.Latitude, req.Longitude, req.Page)
	if err != nil {
		httpCode := http.StatusInternalServerError
		res.Message = err.Error()
		c.JSON(httpCode, res)
		return
	}

	res.Posts = posts
	c.JSON(http.StatusOK, res)
}

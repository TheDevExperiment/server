package models

import "github.com/TheDevExperiment/server/internal/db/posts/models"

type CreatePostRequest struct {
	Post models.Post `json:"post" binding:"required"`
}

type CreatePostResponse struct {
	Message string `form:"message" json:"message" binding:"required"`
}

type NearbyPostRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Page      int     ` json:"page"`
	// Radius int -> if app wants to support tweakable filter
}

type NearbyPostResponse struct {
	Posts   []models.Post `json:"posts" binding:"required"`
	Message string
}

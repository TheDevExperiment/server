package auth

import (
	userModel "github.com/TheDevExperiment/server/internal/db/users/models"
)

type AuthRequest struct {
	SecretToken string `form:"secretToken" json:"secretToken" xml:"secretToken" binding:"required"`
}

type AuthResponse struct {
	Message string           `form:"message" json:"message" xml:"message"  binding:"required"`
	Data    []userModel.User `form:"data" json:"data"`
}

type CreateAccountRequest struct {
	Age       string `bson:"age" json:"age" binding:"required"`
	CountryId string `bson:"countryId" json:"countryId" binding:"required"`
	CityId    string `bson:"cityId" json:"cityId" binding:"required"`
}

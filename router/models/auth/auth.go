package auth

import (
	"github.com/TheDevExperiment/server/internal/db/repositories"
)

type AuthRequest struct {
	SecretToken string `form:"secretToken" json:"secretToken" xml:"secretToken" binding:"required"`
}

type AuthResponse struct {
	Message   string                   `form:"message" json:"message" xml:"message"  binding:"required"`
	ErrorCode string                   `form:"errorCode" json:"errorCode" xml:"errorCode" binding:"required"`
	Data      []repositories.UserModel `form:"data" json:"data"`
}

type CreateAccountRequest struct {
	Age       string `bson:"age" json:"age" binding:"required"`
	CountryId string `bson:"countryId" json:"countryId" binding:"required"`
	CityId    string `bson:"cityId" json:"cityId" binding:"required"`
}

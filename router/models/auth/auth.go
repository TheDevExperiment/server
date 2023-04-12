package auth

import (
	userModel "github.com/TheDevExperiment/server/internal/db/users/models"
)

// gin dereferences the pointer and marshals the struct into JSON for the HTTP response
type Response struct {
	Message string          `form:"message" json:"message" xml:"message"  binding:"required"`
	Data    *userModel.User `form:"data" json:"data"`
}

type CreateAccountRequest struct {
	Age       string `bson:"age" json:"age" binding:"required"`
	CountryId string `bson:"countryId" json:"countryId" binding:"required"`
	CityId    string `bson:"cityId" json:"cityId" binding:"required"`
}
type CreateAccountResponse struct {
	Message string          `form:"message" json:"message" xml:"message"  binding:"required"`
	Data    *userModel.User `form:"data" json:"data"`
}

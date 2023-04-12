package auth

import (
	userModel "github.com/TheDevExperiment/server/internal/db/users/models"
)

type AuthResponse struct {
	Message string `form:"message" json:"message" xml:"message"  binding:"required"`

	//gin dereferences the pointer and marshals the struct into JSON for the HTTP response
	//more here: https://github.com/gin-gonic/gin#json
	Data *userModel.User `form:"data" json:"data"`
}

type CreateAccountRequest struct {
	Age       string `bson:"age" json:"age" binding:"required"`
	CountryId string `bson:"countryId" json:"countryId" binding:"required"`
	CityId    string `bson:"cityId" json:"cityId" binding:"required"`
}

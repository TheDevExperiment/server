package models

import (
	"time"
)

type Score struct {
	RatingScoreSum int `bson:"ratingScoreSum"`
	RatingCount    int `bson:"ratingCount"`
}
type User struct {
	Id              string    `bson:"_id,omitempty" `
	IsGuest         bool      `bson:"isGuest,omitempty"`
	GuestAuthSecret string    `bson:"guestAuthSecret,omitempty"`
	CreatedAt       time.Time `bson:"createdAt,omitempty"`
	LastModified    time.Time `bson:"lastModified,omitempty"`
	DisplayName     string    `bson:"displayName,omitempty"`
	CountryID       string    `bson:"countryId,omitempty"`
	CityID          string    `bson:"cityId,omitempty"`
	Age             string    `bson:"age,omitempty"`
	Score           Score     `bson:"score,omitempty"`
	IsActive        bool      `bson:"isActive,omitempty"`
	DeletionReason  string    `bson:"deletionReason,omitempty"`
}
type UserUpdateModel struct {
	IsGuest         bool      `bson:"isGuest,omitempty"`
	GuestAuthSecret string    `bson:"guestAuthSecret,omitempty"`
	LastModified    time.Time `bson:"lastModified,omitempty"`
	DisplayName     string    `bson:"displayName,omitempty"`
	Score           Score     `bson:"score,omitempty"`
	IsActive        bool      `bson:"isActive,omitempty"`
	DeletionReason  string    `bson:"deletionReason,omitempty"`
}

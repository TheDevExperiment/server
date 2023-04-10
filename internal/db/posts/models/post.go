package models

import "time"

// this struct will be used in internal packages
// for calling Post interface methods.
type Post struct {
	Id            string    `bson:"_id,omitempty"`
	UserId        int       `bson:"userId,omitempty"`
	AuthorName    string    `bson:"authorName,omitempty"`
	CreatedAt     time.Time `bson:"createdAt,omitempty"`
	CityId        int       `bson:"cityId,omitempty"`
	CountryId     int       `bson:"countryId,omitempty"`
	Coordinates   Point     `bson:"coordinates,omitempty"`
	IsActive      bool      `bson:"isActive,omitempty"`
	LikesCount    int       `bson:"likesCount,omitempty"`
	CommentsCount int       `bson:"commentsCount,omitempty"`
	Content       Content   `bson:"content,omitempty,inline"`
}

type Point struct {
	Type        string    `bson:"type"`        //needs to be set to Point
	Coordinates []float64 `bson:"coordinates"` //[long,lat]
}

type Content struct {
	Text           string         `bson:"text,omitempty"`
	BackgroundUrl  string         `bson:"backgroundUrl,omitempty"`
	BackgroundType BackgroundType `bson:"backgroundType"`
}

type BackgroundType int

const (
	PreloadedImage BackgroundType = iota
	PreloadedVideo
	CustomImage
	CustomVideo
)

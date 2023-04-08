package models

import "time"

// this struct will be used in internal packages
// for calling Post interface methods.
type Post struct {
	Id          string
	UserId      string
	AuthorName  string
	CreatedAt   time.Time
	CityId      int
	CountryId   int
	Coordinates Coordinates
	IsActive    bool
	Likes       int
}

//TODO: create mongodb schema which will be private to this pkg

type Coordinates struct {
	latitude, longitude float64
}

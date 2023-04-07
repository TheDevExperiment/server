// internal/repositories/user_repository.go

package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TheDevExperiment/server/internal/db"
	"github.com/TheDevExperiment/server/internal/utility/funnynamer"
)

type UserRepository struct {
	collection *mongo.Collection
}
type UserModel struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" `
	IsGuest         bool               `bson:"isGuest,omitempty"`
	GuestAuthSecret string             `bson:"guestAuthSecret,omitempty"`
	CreatedAt       primitive.DateTime `bson:"createdAt,omitempty"`
	LastModified    primitive.DateTime `bson:"lastModified,omitempty"`
	DisplayName     string             `bson:"displayName,omitempty"`
	CountryID       string             `bson:"countryId,omitempty"`
	CityID          string             `bson:"cityId,omitempty"`
	Age             string             `bson:"age,omitempty"`
	Score           _score             `bson:"score,omitempty"`
	IsActive        bool               `bson:"isActive,omitempty"`
	DeletionReason  string             `bson:"deletionReason,omitempty"`
}

type _score struct {
	RatingScoreSum int `bson:"ratingScoreSum"`
	RatingCount    int `bson:"ratingCount"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db.GetCollection(viper.GetString("mongodb.db_name"), "users")}
}

func (r *UserRepository) Find(ctx context.Context, filter interface{}) ([]UserModel, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []UserModel
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *UserRepository) Create(ctx context.Context, userAge string, countryId string, cityId string) (UserModel, error) {
	uuid := uuid.New().String()
	doc := UserModel{
		IsGuest:         true,
		GuestAuthSecret: uuid,
		CreatedAt:       primitive.NewDateTimeFromTime(time.Now()),
		LastModified:    primitive.NewDateTimeFromTime(time.Now()),
		DisplayName:     funnynamer.GetFunnyName(),
		CountryID:       countryId,
		CityID:          cityId,
		Age:             userAge,
		Score:           _score{RatingScoreSum: 0, RatingCount: 0},
		IsActive:        true,
		DeletionReason:  "",
	}
	result, err := r.collection.InsertOne(context.Background(), doc)
	if err != nil {
		return UserModel{}, err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return UserModel{}, errors.New("Failed to convert InsertedID to ObjectID")
	}
	doc.Id = insertedID
	return doc, nil
}

func (r *UserRepository) Delete(ctx context.Context, filter interface{}) error {
	_, err := r.collection.DeleteMany(ctx, filter)
	return err
}

func (r *UserRepository) Update(ctx context.Context, filter interface{}, update interface{}) error {
	updateDoc := bson.M{
		"$set": update,
	}

	_, err := r.collection.UpdateMany(ctx, filter, updateDoc)
	return err
}

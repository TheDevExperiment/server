// internal/repositories/user_repository.go

package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TheDevExperiment/server/internal/db"
	"github.com/TheDevExperiment/server/internal/log"
	"github.com/TheDevExperiment/server/internal/utility/funnynamer"
	"github.com/TheDevExperiment/server/internal/utility/jwt"
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
type UserUpdateModel struct {
	// Id              primitive.ObjectID `bson:"_id,omitempty" `
	IsGuest         bool               `bson:"isGuest,omitempty"`
	GuestAuthSecret string             `bson:"guestAuthSecret,omitempty"`
	LastModified    primitive.DateTime `bson:"lastModified,omitempty"`
	DisplayName     string             `bson:"displayName,omitempty"`
	// CountryID       string             `bson:"countryId,omitempty"`
	// CityID          string             `bson:"cityId,omitempty"`
	// Age            string `bson:"age,omitempty"`
	Score          _score `bson:"score,omitempty"`
	IsActive       bool   `bson:"isActive,omitempty"`
	DeletionReason string `bson:"deletionReason,omitempty"`
}
type _score struct {
	RatingScoreSum int `bson:"ratingScoreSum"`
	RatingCount    int `bson:"ratingCount"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db.GetCollection(viper.GetString("mongodb.db_name"), db.CollectionUsers)}
}
func (r *UserRepository) FindById(ctx context.Context, id string) (*UserModel, error) {
	bsonFilter := bson.M{db.FieldId: id}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	bsonFilter[db.FieldId] = oid

	result := r.collection.FindOne(ctx, bsonFilter)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, result.Err()
	}
	var user UserModel
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, userAge string, countryId string, cityId string) (UserModel, error) {
	doc := UserModel{
		IsGuest:         true,
		GuestAuthSecret: "",
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

	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return UserModel{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return UserModel{}, errors.New("failed to convert InsertedID to ObjectID")
	}
	secret, err := jwt.CreateToken(insertedID.Hex(), countryId, cityId, true)
	if err != nil {
		return UserModel{}, err
	}

	deltaUpdate := UserUpdateModel{
		GuestAuthSecret: secret,
		LastModified:    primitive.NewDateTimeFromTime(time.Now()),
	}
	success, updateErr := r.UpdateById(ctx, insertedID.Hex(), deltaUpdate)
	if updateErr != nil {
		return UserModel{}, updateErr
	}
	if !success {
		return UserModel{}, nil
	}

	doc.Id = insertedID
	doc.GuestAuthSecret = secret
	doc.LastModified = primitive.NewDateTimeFromTime(time.Now())
	doc.Score.RatingCount = 0
	doc.Score.RatingScoreSum = 0
	doc.IsActive = true
	doc.DeletionReason = ""
	return doc, nil
}

func (r *UserRepository) Delete(ctx context.Context, filter interface{}) error {
	_, err := r.collection.DeleteMany(ctx, filter)
	return err
}

func (r *UserRepository) UpdateById(ctx context.Context, id string, update UserUpdateModel) (bool, error) {
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("failed to convert id to object id: %w", err)
		return false, err
	}

	filter := bson.M{db.FieldId: docId}
	updateDoc := bson.M{
		"$set": update,
	}

	updateResult, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return false, err
	}
	log.Debug(updateResult)
	isSuccess := updateResult.ModifiedCount > 0
	return isSuccess, nil
}

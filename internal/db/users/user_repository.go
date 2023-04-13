// internal/repositories/user_repository.go

package users

import (
	"context"
	"errors"
	"time"

	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TheDevExperiment/server/internal/db"
	"github.com/TheDevExperiment/server/internal/db/users/models"
	"github.com/TheDevExperiment/server/internal/log"
	"github.com/TheDevExperiment/server/internal/utility/funnynamer"
	"github.com/TheDevExperiment/server/internal/utility/jwt"
)

type UserRepository struct {
	collection *mongo.Collection
}

var singletonUserCollection *mongo.Collection

func NewUserRepository() Users {
	if singletonUserCollection == nil {
		log.Debug("Created Connection to user collection")
		singletonUserCollection = db.GetCollection(viper.GetString("mongodb.db_name"), db.CollectionUsers)
	}
	return &UserRepository{collection: singletonUserCollection}
}

func (r *UserRepository) FindById(ctx context.Context, id string) (*models.User, error) {
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
	var user models.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, userAge string, countryId string, cityId string) (*models.User, error) {
	doc := models.User{
		IsGuest:         true,
		GuestAuthSecret: "",
		CreatedAt:       time.Now(),
		LastModified:    time.Now(),
		DisplayName:     funnynamer.GetFunnyName(),
		CountryID:       countryId,
		CityID:          cityId,
		Age:             userAge,
		Score:           models.Score{RatingScoreSum: 0, RatingCount: 0},
		IsActive:        true,
		DeletionReason:  "",
	}

	result, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return &models.User{}, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return &models.User{}, errors.New("failed to convert InsertedID to ObjectID")
	}
	secret, err := jwt.CreateToken(insertedID.Hex(), countryId, cityId, true)
	if err != nil {
		return &models.User{}, err
	}

	deltaUpdate := models.UserUpdateModel{
		GuestAuthSecret: secret,
		LastModified:    time.Now(),
	}
	updateErr := r.UpdateById(ctx, insertedID.Hex(), deltaUpdate)
	if updateErr != nil {
		return &models.User{}, updateErr
	}
	doc.Id = insertedID.Hex()
	doc.GuestAuthSecret = secret
	doc.LastModified = time.Now()
	doc.Score.RatingCount = 0
	doc.Score.RatingScoreSum = 0
	doc.IsActive = true
	doc.DeletionReason = ""
	return &doc, nil
}

func (r *UserRepository) Delete(ctx context.Context, filter interface{}) error {
	_, err := r.collection.DeleteMany(ctx, filter)
	return err
}

func (r *UserRepository) UpdateById(ctx context.Context, id string, update models.UserUpdateModel) error {
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("failed to convert id to object id: %w", err)
		return err
	}

	filter := bson.M{db.FieldId: docId}
	updateDoc := bson.M{
		"$set": update,
	}

	updateResult, err := r.collection.UpdateOne(ctx, filter, updateDoc)
	if err != nil {
		return err
	}
	log.Debug(updateResult)
	isSuccess := updateResult.ModifiedCount > 0
	if !isSuccess {
		return fmt.Errorf("unsuccesful insert")
	}
	return nil
}

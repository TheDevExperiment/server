// internal/repositories/user_repository.go

package repositories

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/TheDevExperiment/server/internal/db"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db.GetCollection(viper.GetString("mongodb.db_name"), "users")}
}

func (r *UserRepository) Find(ctx context.Context, filter interface{}) ([]interface{}, error) {
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	fmt.Println("Result = ")
	fmt.Println(results)
	return results, nil
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

func (r *UserRepository) GetById(ctx context.Context, id primitive.ObjectID) (interface{}, error) {
	filter := bson.M{"_id": id}
	result := r.collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return nil, err
	}

	var user interface{}
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

package posts

import (
	"context"

	"github.com/TheDevExperiment/server/db/posts/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostsRepository struct {
	collection *mongo.Collection
}

func NewPostsRepository() *PostsRepository {
	// use mongodb to find collection here
	return nil
}

func (pr *PostsRepository) Create(context.Context, models.Post) error {
	return nil
}

func (pr *PostsRepository) FindNearby(context.Context, models.Coordinates) ([]models.Post, error) {
	return []models.Post{}, nil
}

func (pr *PostsRepository) Like(context.Context, string) error {
	return nil
}

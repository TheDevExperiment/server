package posts

import (
	"context"

	"github.com/TheDevExperiment/server/internal/db"
	"github.com/TheDevExperiment/server/internal/db/posts/models"
	"github.com/TheDevExperiment/server/internal/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostsRepository struct {
	collection *mongo.Collection
}

var singletonPostCollection *mongo.Collection

const defaultNearbySearchRadius = 1000 // 1km
const defaultPageLimit = 10

func NewPostsRepository() *PostsRepository {
	// use mongodb to find collection here
	if singletonPostCollection == nil {
		log.Debug("Created connection to post collection")
		singletonPostCollection = db.GetCollection(viper.GetString("mongodb.db_name"), "posts")
	}
	return &PostsRepository{collection: singletonPostCollection}
}

func (pr *PostsRepository) Create(ctx context.Context, post models.Post) error {
	log.Debugf("Received new post create request data: %+v", post)
	res, err := pr.collection.InsertOne(ctx, post)
	log.Debugf("Response received from mongo: %+v", res)
	return err
}

func prepareNearbySortQuery(latitude, longitude float64) primitive.M {
	query := bson.M{
		"coordinates": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
				"$maxDistance": defaultNearbySearchRadius,
			},
		},
	}
	return query
}

func preparePaginatedNearbySortOpts(page int) *options.FindOptions {
	findOptions := options.Find()
	findOptions.SetLimit(int64(defaultPageLimit))             // Set the limit to the specified value
	findOptions.SetSkip(int64((page - 1) * defaultPageLimit)) // Skip records based on the page number and limit
	findOptions.SetSort(bson.D{{Key: "_id", Value: -1}})      // Sort by "_id" field in descending order
	findOptions.SetProjection(bson.M{"coordinates": 0})       // Exclude "coordinates" field from the result
	return findOptions
}

func (pr *PostsRepository) FindNearby(ctx context.Context, lat, long float64, page int) ([]models.Post, error) {
	log.Debugf("Received new nearby req data: %f %f", lat, long)
	query := prepareNearbySortQuery(lat, long)
	optsWithPagination := preparePaginatedNearbySortOpts(page)

	cur, err := pr.collection.Find(ctx, query, optsWithPagination)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var posts []models.Post
	for cur.Next(ctx) {
		var post models.Post
		err := cur.Decode(&post)
		if err != nil {
			log.Errorf("Error while parsing doc: %s", err.Error())
			continue
		}
		posts = append(posts, post)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *PostsRepository) Like(context.Context, string) error {
	return nil
}

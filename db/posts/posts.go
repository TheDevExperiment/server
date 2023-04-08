package posts

import (
	"context"

	"github.com/TheDevExperiment/server/db/posts/models"
)

type Posts interface {
	Create(context.Context, models.Post) error
	FindNearby(context.Context, models.Coordinates) ([]models.Post, error)
	Like(context.Context, string) error
}

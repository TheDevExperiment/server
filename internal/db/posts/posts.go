package posts

import (
	"context"

	"github.com/TheDevExperiment/server/internal/db/posts/models"
)

type Posts interface {
	Create(context.Context, models.Post) error
	FindNearby(context.Context, float64, float64) ([]models.Post, error)
	Like(context.Context, string) error
}

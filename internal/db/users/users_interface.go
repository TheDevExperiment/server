package users

import (
	"context"

	"github.com/TheDevExperiment/server/internal/db/users/models"
)

type Users interface {
	FindById(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, userAge string, countryId string, cityId string) (*models.User, error)
	Delete(ctx context.Context, filter interface{}) error
	UpdateById(ctx context.Context, id string, update models.UserUpdateModel) (bool, error)
}

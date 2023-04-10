package repositories

import (
	"context"
)

type Repository interface {
	Find(ctx context.Context, filter interface{}) ([]interface{}, error)
	Delete(ctx context.Context, filter interface{}) error
	Update(ctx context.Context, filter interface{}, update interface{}) error
}

package auth

import (
	"context"

	"github.com/TheDevExperiment/server/router/models"
)

func Handler(ctx context.Context, req *models.AuthRequest) (*models.AuthResponse, error) {
	response := &models.AuthResponse{}
	response.Msg = "invalid creds"

	if req.User == "test" && req.Password == "test" {
		response.Msg = "welcome test user"
	}

	return response, nil
}

package usecacse

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
	"github.com/ngobrut/eniqlo-store-api/internal/types/request"
	"github.com/ngobrut/eniqlo-store-api/internal/types/response"
)

type IFaceUsecase interface {
	// auth
	Register(ctx context.Context, req *request.Register) (*response.AuthResponse, error)
	Login(ctx context.Context, req *request.Login) (*response.AuthResponse, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*model.User, error)
}

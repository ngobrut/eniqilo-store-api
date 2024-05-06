package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/internal/model"
)

type IFaceRepository interface {
	// user
	CreateUser(ctx context.Context, data *model.User) error
	GetOneUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetOneUserByEmail(ctx context.Context, email string) (*model.User, error)

	// product
	CreateProduct(ctx context.Context, data *model.Product) error
}

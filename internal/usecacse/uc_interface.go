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

	// product
	CreateProduct(ctx context.Context, req *request.CreateProduct) (*response.CreateProduct, error)
	GetListProduct(ctx context.Context, req *request.ListProductQuery) ([]*response.ListProduct, error)
	UpdateProduct(ctx context.Context, req *request.UpdateProduct) error
	DeleteProduct(ctx context.Context, productID uuid.UUID) error

	// search sku
	SearchSKU(ctx context.Context, req *request.SearchQuery) ([]*response.SearchSKU, error)

	// customer
	RegisterCustomer(ctx context.Context, req *request.RegisterCustomer) (*response.RegisterCustomer, error)
}

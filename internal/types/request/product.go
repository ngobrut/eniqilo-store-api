package request

import (
	"github.com/google/uuid"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
)

type CreateProduct struct {
	Name        string            `json:"name" validate:"required,min=1,max=30"`
	Sku         string            `json:"sku" validate:"required,min=1,max=30"`
	Category    constant.Category `json:"category" validate:"required,category"`
	ImageUrl    string            `josn:"imageUrl" validate:"required,validUrl"`
	Notes       string            `json:"notes" validate:"required,min=1,max=200"`
	Price       int               `json:"price" validate:"required,min=1"`
	Stock       *int              `json:"stock" validate:"required,min=0,max=100000"`
	Location    string            `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool             `json:"isAvailable" validate:"required"`
}

type ListProductQuery struct {
	ID          *string
	Limit       *int
	Offset      *int
	Name        *string
	IsAvailable *bool
	Category    *string
	Sku         *string
	InStock     *bool
	Price       *string
	CreatedAt   *string
}

type UpdateProduct struct {
	Name        string            `json:"name" validate:"required,min=1,max=30"`
	Sku         string            `json:"sku" validate:"required,min=1,max=30"`
	Category    constant.Category `json:"category" validate:"required,category"`
	ImageUrl    string            `josn:"imageUrl" validate:"required,validUrl"`
	Notes       string            `json:"notes" validate:"required,min=1,max=200"`
	Price       int               `json:"price" validate:"required,min=1"`
	Stock       *int              `json:"stock" validate:"required,min=0,max=100000"`
	Location    string            `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool             `json:"isAvailable" validate:"required"`
	ProductID   uuid.UUID         `json:"-"`
}

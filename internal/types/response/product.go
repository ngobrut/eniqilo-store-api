package response

import (
	"github.com/google/uuid"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
)

type CreateProduct struct {
	ProductID string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ListProduct struct {
	ProductID   uuid.UUID         `json:"id" db:"product_id"`
	Name        string            `json:"name" db:"name"`
	Sku         string            `json:"sku" db:"sku"`
	Category    constant.Category `json:"category" db:"category"`
	ImageUrl    string            `json:"imageUrl" db:"image_url"`
	Stock       int               `json:"stock" db:"stock"`
	Notes       string            `json:"notes" db:"notes"`
	Price       int               `json:"price" db:"price"`
	Location    string            `json:"location" db:"location"`
	IsAvailable bool              `json:"isAvailable" db:"is_available"`
	CreatedAt   string            `json:"createdAt" db:"created_at"`
}

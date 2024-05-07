package response

import (
	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
)

type CreateProduct struct {
	ProductID string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ListProduct struct {
	ProductID   uuid.UUID         `json:"product_id" db:"product_id"`
	Name        string            `json:"name" db:"name"`
	Sku         string            `json:"sku" db:"sku"`
	Category    constant.Category `json:"category" db:"category"`
	ImageUrl    string            `josn:"imageUrl" db:"image_url"`
	Stock       int               `json:"stock" db:"stock"`
	Notes       string            `json:"notes" db:"notes"`
	Price       int               `json:"price" db:"price"`
	Location    string            `json:"location" db:"location"`
	IsAvailable bool              `json:"isAvailable" db:"is_available"`
	CreatedAt   string            `json:"created_at" db:"created_at"`
}

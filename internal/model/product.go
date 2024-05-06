package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngobrut/eniqlo-store-api/pkg/constant"
)

type Product struct {
	ProductID   uuid.UUID         `json:"product_id" db:"product_id"`
	Name        string            `json:"name" db:"name"`
	Sku         string            `json:"sku" db:"sku"`
	Category    constant.Category `json:"category" db:"category"`
	ImageUrl    string            `josn:"imageUrl" db:"image_url"`
	Notes       string            `json:"notes" db:"notes"`
	Price       int               `json:"price" db:"price"`
	Stock       int               `json:"stock" db:"stock"`
	Location    string            `json:"location" db:"location"`
	IsAvailable bool              `json:"isAvailable" db:"is_available"`
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time        `json:"deleted_at" db:"deleted_at"`
}

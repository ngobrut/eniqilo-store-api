package response

import (
	"github.com/google/uuid"
	"github.com/ngobrut/eniqilo-store-api/pkg/constant"
)

type SearchSKU struct {
	ProductID uuid.UUID         `json:"product_id" db:"product_id"`
	Name      string            `json:"name" db:"name"`
	Sku       string            `json:"sku" db:"sku"`
	Category  constant.Category `json:"category" db:"category"`
	ImageUrl  string            `josn:"imageUrl" db:"image_url"`
	Stock     int               `json:"stock" db:"stock"`
	Price     int               `json:"price" db:"price"`
	Location  string            `json:"location" db:"location"`
	CreatedAt string            `json:"created_at" db:"created_at"`
}

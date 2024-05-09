package model

import (
	"time"

	"github.com/google/uuid"
)

type InvoiceProduct struct {
	ID        int       `json:"-" db:"id"`
	InvoiceID uuid.UUID `json:"transactionId" db:"invoice_id"`
	ProductID uuid.UUID `json:"productId" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
	Price     int       `json:"price" db:"price"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

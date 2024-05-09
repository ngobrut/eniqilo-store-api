package model

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	InvoiceID  uuid.UUID `json:"transactionId" db:"invoice_id"`
	CustomerID uuid.UUID `json:"customerId" db:"customer_id"`
	TotalPrice int       `json:"-" db:"total_price"`
	Paid       int       `json:"paid" db:"paid"`
	Change     int       `json:"change" db:"change"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

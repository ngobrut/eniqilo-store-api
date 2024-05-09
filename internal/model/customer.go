package model

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerID uuid.UUID `json:"userId" db:"customer_id"`
	Name       string    `json:"name" db:"name"`
	Phone      string    `json:"phoneNumber" db:"phone"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

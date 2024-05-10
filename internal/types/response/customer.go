package response

import "github.com/google/uuid"

type RegisterCustomer struct {
	CustomerID uuid.UUID `json:"userId" db:"customer_id"`
	Phone      string    `json:"phoneNumber" db:"phone"`
	Name       string    `json:"name" db:"name"`
}

type ListCustomer struct {
	CustomerID uuid.UUID `json:"userId" db:"customer_id"`
	Phone      string    `json:"phoneNumber" db:"phone"`
	Name       string    `json:"name" db:"name"`
}

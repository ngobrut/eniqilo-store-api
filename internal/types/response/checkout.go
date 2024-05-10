package response

import "github.com/google/uuid"

type ProductDetail struct {
	ProductID uuid.UUID `json:"productId" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
}

type ListInvoice struct {
	InvoiceID     uuid.UUID       `json:"transactionId" db:"invoice_id"`
	CustomerID    uuid.UUID       `json:"customerId" db:"customer_id"`
	ProductDetail []ProductDetail `json:"productDetails" db:"-"`
	Paid          int             `json:"paid" db:"paid"`
	Change        int             `json:"change" db:"change"`
	CreatedAt     string          `json:"createdAt" db:"created_at"`
}

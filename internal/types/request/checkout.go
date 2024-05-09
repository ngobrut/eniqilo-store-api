package request

type ProductCheckout struct {
	ProductID string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type Checkout struct {
	CustomerID     string            `json:"customerId" validate:"required"`
	ProductDetails []ProductCheckout `json:"productDetails" validate:"required,min=1,dive"`
	Paid           int               `json:"paid" validate:"required,min=1"`
	Change         *int              `josn:"change" validate:"required,min=0"`
}

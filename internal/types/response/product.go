package response

import "time"

type CreateProduct struct {
	ProductID string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

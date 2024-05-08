package response

import "github.com/google/uuid"

type AuthResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	Phone       string    `json:"phone"`
	Name        string    `json:"name"`
	AccessToken string    `json:"accessToken,omitempty"`
}

package response

import "github.com/google/uuid"

type AuthResponse struct {
	UserID      uuid.UUID `json:"userId"`
	Phone       string    `json:"phoneNumber"`
	Name        string    `json:"name"`
	AccessToken string    `json:"accessToken,omitempty"`
}

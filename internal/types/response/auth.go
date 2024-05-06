package response

import "time"

type AuthResponse struct {
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	AccessToken string    `json:"accessToken,omitempty"`
}

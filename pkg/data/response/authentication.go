package response

import (
	"time"
)


type UsersResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
}


type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}

type Logout struct {
	Status string
}

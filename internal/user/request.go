package user

import "time"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SaveRefreshTokenRequest struct {
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type GetRefreshTokenResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

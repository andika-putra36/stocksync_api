package user

type GetLoginCredentialResponse struct {
	UserID       int    `json:"user_id"`
	RoleID       int    `json:"role_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

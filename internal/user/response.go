package user

type GetLoginCredentialResponse struct {
	RoleID       string `json:"role_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type LoginResponse struct {
	IsLoggedIn bool `json:"is_logged_in"`
}

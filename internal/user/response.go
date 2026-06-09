package user

type GetUserInformationResponse struct {
	UserName string `json:"user_name"`
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	IsActive bool   `json:"is_active"`
}

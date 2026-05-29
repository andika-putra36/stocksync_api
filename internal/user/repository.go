package user

import "gorm.io/gorm"

type Repository interface {
	GetLoginCredentials(input LoginRequest) (GetLoginCredentialResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetLoginCredentials(input LoginRequest) (GetLoginCredentialResponse, error) {
	var response GetLoginCredentialResponse

	err := r.db.Raw(
		`SELECT * FROM get_login_credential(?)`,
		input.Email,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

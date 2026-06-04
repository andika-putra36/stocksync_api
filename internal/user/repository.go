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

func (r *repository) SaveRefreshToken(input SaveRefreshTokenRequest) error {
	err := r.db.Exec(
		`CALL save_refresh_token(?, ?, ?)`,
		input.UserID,
		input.Token,
		input.ExpiredAt,
	).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetRefreshToken(token string) (GetRefreshTokenResponse, error) {
	var response GetRefreshTokenResponse

	err := r.db.Raw(
		`SELECT * FROM get_refresh_token(?)`,
		token,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

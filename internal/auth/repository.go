package auth

import "gorm.io/gorm"

type Repository interface {
	GetLoginCredentials(input LoginRequest) (GetLoginCredentialResponse, error)
	SaveRefreshToken(input SaveRefreshTokenRequest) error
	GetRefreshToken(refreshToken string) (GetRefreshTokenResponse, error)
	DeleteRefreshToken(userID int) error
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

func (r *repository) GetRefreshToken(refreshToken string) (GetRefreshTokenResponse, error) {
	var response GetRefreshTokenResponse

	err := r.db.Raw(
		`SELECT * FROM get_refresh_token(?)`,
		refreshToken,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

func (r *repository) DeleteRefreshToken(userID int) error {
	err := r.db.Exec(
		`CALL delete_refresh_token(?)`,
		userID,
	).Error
	if err != nil {
		return err
	}
	return nil
}

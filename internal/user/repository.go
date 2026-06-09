package user

import "gorm.io/gorm"

type Repository interface {
	GetUserInformation(userID int) (GetUserInformationResponse, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUserInformation(userID int) (GetUserInformationResponse, error) {
	var response GetUserInformationResponse

	err := r.db.Raw(
		`SELECT * FROM get_user_information(?)`,
		userID,
	).Scan(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

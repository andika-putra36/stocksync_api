package user

import (
	"stocksync_api/pkg/bcrypt"
	"stocksync_api/pkg/jwt"
)

type Service interface {
	LogIn(input LoginRequest) (LoginResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) LogIn(input LoginRequest) (LoginResponse, error) {
	var loginCredential GetLoginCredentialResponse

	// Fetch data
	loginCredential, err := s.repository.GetLoginCredentials(input)
	if err != nil {
		return LoginResponse{}, err
	}

	// Compare password
	err = bcrypt.ComparePassword(loginCredential.PasswordHash, input.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	// Generate access token
	accessToken, err := jwt.GenerateAccessToken(loginCredential.UserID, loginCredential.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		// IsLoggedIn:  true,
		AccessToken: accessToken,
	}, nil
}

// func ComparePassword(hashedPassword, plainPassword string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
// 	if err != nil {
// 		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
// 			return errors.New("invalid credentials")
// 		}
// 		return errors.New("failed to compare password")
// 	}
// 	return nil
// }

package auth

import (
	"errors"
	"stocksync_api/internal/user"
	"stocksync_api/pkg/bcrypt"
	"stocksync_api/pkg/jwt"
	"time"
)

type Service interface {
	LogIn(input LoginRequest) (LoginResponse, error)
	RefreshToken(input RefreshTokenRequest) (RefreshTokenResponse, error)
}

type service struct {
	repository     Repository
	userRepository user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *service {
	return &service{
		repository,
		userRepository,
	}
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

	// Generate refresh token
	refreshToken, expiredAt, err := jwt.GenerateRefreshToken()
	if err != nil {
		return LoginResponse{}, err
	}

	// Delete refresh token in DB if exists
	err = s.repository.DeleteRefreshToken(loginCredential.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	// Store refresh token in DB
	err = s.repository.SaveRefreshToken(SaveRefreshTokenRequest{
		UserID:    loginCredential.UserID,
		Token:     refreshToken,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	userInformation, err := s.userRepository.GetUserInformation(loginCredential.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		// IsLoggedIn:  true,
		User: user.GetUserInformationResponse{
			UserName: userInformation.UserName,
			RoleID:   userInformation.RoleID,
			RoleName: userInformation.RoleName,
			IsActive: userInformation.IsActive,
		},
		Token: Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (s *service) RefreshToken(input RefreshTokenRequest) (RefreshTokenResponse, error) {
	tokenData, err := s.repository.GetRefreshToken(input.RefreshToken)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	if time.Now().UTC().After(tokenData.ExpiredAt) {
		return RefreshTokenResponse{}, errors.New("Refresh token expired")
	}

	accessToken, err := jwt.GenerateAccessToken(tokenData.UserID, tokenData.Email)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	newRefreshToken, expiredAt, err := jwt.GenerateRefreshToken()
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	err = s.repository.DeleteRefreshToken(tokenData.UserID)
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	err = s.repository.SaveRefreshToken(SaveRefreshTokenRequest{
		UserID:    tokenData.UserID,
		Token:     newRefreshToken,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return RefreshTokenResponse{}, err
	}

	return RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

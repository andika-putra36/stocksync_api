package auth

import (
	"errors"
	"stocksync_api/pkg/bcrypt"
	"stocksync_api/pkg/jwt"
	"time"
)

type Service interface {
	LogIn(input LoginRequest) (LoginResponse, error)
	RefreshToken(input RefreshTokenRequest) (LoginResponse, error)
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

	// Generate refresh token
	refreshToken, expiredAt, err := jwt.GenerateRefreshToken()
	if err != nil {
		return LoginResponse{}, err
	}

	//Store refresh token in DB
	err = s.repository.SaveRefreshToken(SaveRefreshTokenRequest{
		UserID:    loginCredential.UserID,
		Token:     refreshToken,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		// IsLoggedIn:  true,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) RefreshToken(input RefreshTokenRequest) (LoginResponse, error) {
	tokenData, err := s.repository.GetRefreshToken(input.RefreshToken)
	if err != nil {
		return LoginResponse{}, err
	}

	if time.Now().UTC().After(tokenData.ExpiredAt) {
		return LoginResponse{}, errors.New("Refresh token expired")
	}

	accessToken, err := jwt.GenerateAccessToken(tokenData.UserID, tokenData.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	newRefreshToken, expiredAt, err := jwt.GenerateRefreshToken()
	if err != nil {
		return LoginResponse{}, err
	}

	err = s.repository.DeleteRefreshToken(tokenData.UserID)
	if err != nil {
		return LoginResponse{}, err
	}

	err = s.repository.SaveRefreshToken(SaveRefreshTokenRequest{
		UserID:    tokenData.UserID,
		Token:     newRefreshToken,
		ExpiredAt: expiredAt,
	})
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

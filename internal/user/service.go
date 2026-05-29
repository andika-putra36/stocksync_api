package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
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

	loginCredential, err := s.repository.GetLoginCredentials(input)
	if err != nil {
		return LoginResponse{}, err
	}

	err = ComparePassword(loginCredential.PasswordHash, input.Password)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{IsLoggedIn: true}, nil
}

func ComparePassword(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("invalid credentials")
		}
		return errors.New("failed to compare password")
	}
	return nil
}

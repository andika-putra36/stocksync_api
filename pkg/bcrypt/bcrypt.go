package bcrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

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

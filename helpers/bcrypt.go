package helpers

import (
	"github.com/burntcarrot/quii/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.ErrInternalServerError
	}
	return string(hash), nil
}

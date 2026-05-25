package auth

import (
	"errors"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hashed_password, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		return "", errors.New("Error while hashing password")
	}

	return hashed_password, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {
	same, err := argon2id.ComparePasswordAndHash(password, hash)

	if err != nil {
		return false, errors.New("Error while comparing password and hash")
	}

	return same, nil
}

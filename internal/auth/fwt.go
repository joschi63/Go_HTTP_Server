package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var now = time.Now

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	issuedAt := now()
	claims := &jwt.RegisteredClaims{
		Issuer:    "chirpy-access",
		IssuedAt:  jwt.NewNumericDate(issuedAt),
		ExpiresAt: jwt.NewNumericDate(issuedAt.Add(expiresIn)),
		Subject:   userID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed_token, err := token.SignedString([]byte(tokenSecret))

	if err != nil {
		return "", fmt.Errorf("error occurred while signing jwt token: %w", err)
	}

	return signed_token, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error occurred while unpacking jwt token: %w", err)
	}

	expireDate, err := token.Claims.GetExpirationTime()

	if err != nil || expireDate == nil {
		return uuid.UUID{}, fmt.Errorf("error occurred while unpacking jwt token: %w", err)
	}

	if !expireDate.Time.After(now()) {
		return uuid.UUID{}, fmt.Errorf("jwt token is expired")
	}

	id, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error occurred while unpacking jwt token: %w", err)
	}

	return uuid.MustParse(id), nil
}

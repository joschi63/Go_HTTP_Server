package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	originalNow := now
	fixedNow := time.Date(2026, time.May, 31, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return fixedNow }
	defer func() { now = originalNow }()

	userID := uuid.MustParse("e38e43a4-fb72-4c64-acc8-d1d0e4aa2e77")
	secret := "ejsxXAhIA5Gj+RIxcIYv3oWB3deybY7bTjO0eKfK7RM="

	tokenString, err := MakeJWT(userID, secret, 24*time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		t.Fatalf("unexpected claims type: %T", token.Claims)
	}

	if claims.Issuer != "chirpy-access" {
		t.Fatalf("unexpected issuer: got %q want %q", claims.Issuer, "chirpy-access")
	}

	if claims.Subject != userID.String() {
		t.Fatalf("unexpected subject: got %q want %q", claims.Subject, userID.String())
	}

	if claims.IssuedAt == nil || !claims.IssuedAt.Time.Equal(fixedNow) {
		t.Fatalf("unexpected issued-at: got %v want %v", claims.IssuedAt, fixedNow)
	}

	wantExpiresAt := fixedNow.Add(24 * time.Hour)
	if claims.ExpiresAt == nil || !claims.ExpiresAt.Time.Equal(wantExpiresAt) {
		t.Fatalf("unexpected expires-at: got %v want %v", claims.ExpiresAt, wantExpiresAt)
	}
}

func TestValidateJWT(t *testing.T) {
	originalNow := now
	fixedNow := time.Date(2026, time.May, 31, 12, 0, 0, 0, time.UTC)
	now = func() time.Time { return fixedNow }
	defer func() { now = originalNow }()

	userID := uuid.MustParse("e38e43a4-fb72-4c64-acc8-d1d0e4aa2e77")
	secret := "ejsxXAhIA5Gj+RIxcIYv3oWB3deybY7bTjO0eKfK7RM="

	tokenString, err := MakeJWT(userID, secret, 24*time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT returned error: %v", err)
	}

	gotUserID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateJWT returned error: %v", err)
	}

	if gotUserID != userID {
		t.Fatalf("unexpected user id: got %s want %s", gotUserID, userID)
	}
}

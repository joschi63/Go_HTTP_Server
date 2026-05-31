package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() string {
	token := make([]byte, 32)

	rand.Read(token)

	str_token := hex.EncodeToString(token)

	return str_token
}

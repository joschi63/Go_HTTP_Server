package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	str := headers.Get("Authorization")

	if str == "" {
		return "", fmt.Errorf("No Bearer token in header")
	}

	token := strings.TrimPrefix(str, "Bearer ")

	return token, nil
}

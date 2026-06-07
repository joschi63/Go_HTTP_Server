package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	str := headers.Get("Authorization")

	if str == "" {
		return "", fmt.Errorf("No ApiKey token in header")
	}

	key := strings.TrimPrefix(str, "ApiKey ")

	return key, nil
}

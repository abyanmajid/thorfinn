package security

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateSecureCode() (string, error) {
	bytes := make([]byte, 20)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	code := base32.NewEncoding("ABCDEFGHJKLMNPQRSTUVWXYZ23456789").EncodeToString(bytes)
	return code, nil
}

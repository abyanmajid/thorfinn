package security

import (
	"golang.org/x/crypto/bcrypt"
)

func HashValue(value string) (string, error) {
	hashedValue, err := bcrypt.GenerateFromPassword([]byte(value), 12)
	if err != nil {
		return "", err
	}
	return string(hashedValue), nil
}

func VerifyHash(hashedValue, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(value))
	return err == nil
}

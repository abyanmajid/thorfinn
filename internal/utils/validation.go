package utils

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

func IsValidPassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()-=_+[]{}|;:,.<>?/`~", char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}
	return nil
}

func IsValidEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func IsValidName(name string) error {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return errors.New("name cannot be empty")
	} else if len(name) > 128 {
		return errors.New("name cannot exceed 128 characters")
	}
	for _, char := range name {
		if !unicode.IsLetter(char) && char != ' ' {
			return errors.New("name must only contain letters and spaces")
		}
	}
	return nil
}

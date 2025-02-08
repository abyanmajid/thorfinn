package auth_features

import (
	"errors"

	"github.com/abyanmajid/v"
)

func validateRegisterPayload(payload RegisterRequest) error {
	email := v.String("Email").Email().Parse(payload.Email)
	password := v.String("Password").Min(8).Parse(payload.Password)

	if !email.Ok {
		return errors.New("invalid email")
	}

	if !password.Ok {
		return errors.New("password must be at least 8 characters long")
	}

	if password.Value != payload.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	return nil
}

func validateUserIdInterface(userIdInterface interface{}) (string, error) {
	userId := v.String("UserId").Parse(userIdInterface)

	if !userId.Ok {
		return "", errors.New("invalid user id")
	}

	return userId.Value, nil
}

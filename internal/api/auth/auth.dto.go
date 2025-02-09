package auth_features

import (
	"errors"
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
)

func GenericError[T any]() *ctx.Response[T] {
	return &ctx.Response[T]{
		Error:      errors.New("an error occurred while processing your request"),
		StatusCode: http.StatusInternalServerError,
	}
}

func CustomError[T any](message string) *ctx.Response[T] {
	return &ctx.Response[T]{
		Error:      errors.New(message),
		StatusCode: http.StatusInternalServerError,
	}
}

type RegisterRequest struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type ConfirmEmailRequest struct {
	Token string `json:"token"`
}

type ConfirmEmailResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

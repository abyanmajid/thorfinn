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

type RegisterRequest struct {
	Email           string
	Password        string
	ConfirmPassword string
}

type RegisterResponse struct {
	Message string
}

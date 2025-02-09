package internal

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

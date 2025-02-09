package auth_features

import (
	"errors"
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/thorfinn/internal/database"
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
	Message       string `json:"message"`
	AccessToken   string `json:"access_token"`
	RefreshTokens string `json:"refresh_token"`
}

type LogoutRequest struct{}

type LogoutResponse struct {
	Message string `json:"message"`
}

type SendVerificationEmailRequest struct {
	Email string `json:"email"`
}

type SendVerificationEmailResponse struct {
	Message string `json:"message"`
}

type SendPasswordResetRequest struct {
	Email string `json:"email"`
}

type SendPasswordResetResponse struct {
	Message string `json:"message"`
}

type ResetPasswordRequest struct {
	Token string
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type OtpSendRequest struct {
	Token string `json:"token"`
}

type OtpSendResponse struct {
	Message string `json:"message"`
}

type OtpVerifyRequest struct {
	Token string `json:"token"`
}

type OtpVerifyResponse struct {
	Message string `json:"message"`
}

type GetSelfRequest struct{}

type GetSelfResponse struct {
	Message string `json:"message"`
}

type UpdateSelfRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateSelfResponse struct {
	Message string `json:"message"`
}

type DeleteSelfRequest struct {
	Password string `json:"password"`
}

type DeleteSelfResponse struct {
	Message string `json:"message"`
}

type GetAllUsersRequest struct{}

type GetAllUsersResponse struct {
	Message string                  `json:"message"`
	Users   []database.ThorfinnUser `json:"users"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	Message string `json:"message"`
}

type UpdateUserRequest struct {
	ID string `json:"id"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

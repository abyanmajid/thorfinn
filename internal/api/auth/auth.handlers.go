package auth_features

import (
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/thorfinn/internal"
)

type AuthHandlers struct {
	config *internal.EnvConfig
}

func NewHandlers(config *internal.EnvConfig) *AuthHandlers {
	return &AuthHandlers{
		config: config,
	}
}

func (h *AuthHandlers) Register(c *ctx.Request[RegisterRequest]) *ctx.Response[RegisterResponse] {
	// Step 1 - Early return if user already exists

	// Step 2 - Validate inputs

	// Step 3 - Send confirmation email

	return &ctx.Response[RegisterResponse]{
		Response: RegisterResponse{
			Message: "User registered successfully",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

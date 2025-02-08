package auth_features

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/matcha/security"
	"github.com/abyanmajid/thorfinn/internal"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type AuthHandlers struct {
	config  *internal.EnvConfig
	queries *database.Queries
	mailer  *email.Client
}

func NewHandlers(config *internal.EnvConfig, queries *database.Queries, mailer *email.Client) *AuthHandlers {
	return &AuthHandlers{
		config:  config,
		queries: queries,
		mailer:  mailer,
	}
}

const (
	RegisterSuccessMessage = "If this email is not already in use, you will receive a confirmation email shortly."
)

func (h *AuthHandlers) Register(c *ctx.Request[RegisterRequest]) *ctx.Response[RegisterResponse] {
	// Step 1 - Early return if user already exists
	_, err := h.queries.FindUserByEmail(c.Request.Context(), c.Body.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Debug("Error finding user by email: %v", err)
		return GenericError[RegisterResponse]()
	}

	if errors.Is(err, sql.ErrNoRows) {
		// Step 2 - Validate inputs
		err := validateRegisterPayload(c.Body)
		if err != nil {
			return &ctx.Response[RegisterResponse]{
				StatusCode: http.StatusBadRequest,
				Error:      err,
			}
		}

		// Step 3 - Create and sign JWT
		token := security.NewJWT(security.JwtClaims{
			"sub":  "1234567890",
			"name": c.Body.Email,
			"iat":  time.Now().Unix(),
			"exp":  time.Now().Add(time.Minute * 10).Unix(),
		})

		signedToken, err := token.Sign([]byte(h.config.JwtSecret))
		if err != nil {
			return GenericError[RegisterResponse]()
		}

		encryptedSignedToken, err := security.Encrypt([]byte(signedToken), []byte(h.config.EncryptionSecret), []byte(h.config.EncryptionIv))
		if err != nil {
			logger.Debug("Error encrypting signed token: %v, key: %s", err, h.config.EncryptionSecret)
			return GenericError[RegisterResponse]()
		}

		tokenUrlSafe := security.EncodeBase64(encryptedSignedToken)

		// Step 4 - Send confirmation email
		verificationLink := fmt.Sprintf("%s/auth/verify-email?token=%s", h.config.Origin, tokenUrlSafe)
		err = h.mailer.SendEmail(h.config.EmailFrom, []string{c.Body.Email}, "Email Verification", "email_verification", map[string]any{
			"VerificationLink": verificationLink,
		})

		if err != nil {
			return GenericError[RegisterResponse]()
		}
	}

	return &ctx.Response[RegisterResponse]{
		Response: RegisterResponse{
			Message: RegisterSuccessMessage,
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

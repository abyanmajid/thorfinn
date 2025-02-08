package auth_features

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

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
	_, err := h.queries.FindUserByEmail(c.Request.Context(), c.Body.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Debug("Error finding user by email: %v", err)
		return GenericError[RegisterResponse]()
	}

	if errors.Is(err, sql.ErrNoRows) {
		err := validateRegisterPayload(c.Body)
		if err != nil {
			return CustomError[RegisterResponse](err.Error())
		}

		passwordHash, err := security.Hash([]byte(c.Body.Password))
		if err != nil {
			return GenericError[RegisterResponse]()
		}

		user, err := h.queries.CreateUser(c.Request.Context(), database.CreateUserParams{
			ID:           uuid.New().String(),
			Email:        c.Body.Email,
			PasswordHash: string(passwordHash.Hash),
		})
		if err != nil {
			return GenericError[RegisterResponse]()
		}

		token := security.NewJWT(security.JwtClaims{
			"user_id": user.ID,
			"iat":     time.Now().Unix(),
			"exp":     time.Now().Add(time.Minute * 10).Unix(),
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

		verificationLink := fmt.Sprintf("%s/auth/verify-email?token=%s", h.config.FrontendUrl, tokenUrlSafe)
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
		StatusCode: http.StatusCreated,
		Error:      nil,
	}
}

func (h *AuthHandlers) VerifyEmail(c *ctx.Request[ConfirmEmailRequest]) *ctx.Response[ConfirmEmailResponse] {
	urlSafeToken := c.GetQueryParam("token")
	if urlSafeToken == "" {
		logger.Error("Token not found")
		return CustomError[ConfirmEmailResponse]("token not found")
	}

	encryptedToken, err := security.DecodeBase64(urlSafeToken)
	if err != nil {
		logger.Error("Error decoding token: %v", err)
		return GenericError[ConfirmEmailResponse]()
	}

	tokenByte, err := security.Decrypt([]byte(encryptedToken), []byte(h.config.EncryptionSecret))
	if err != nil {
		logger.Error("Error decrypting token: %v", err)
		return GenericError[ConfirmEmailResponse]()
	}
	verifiedToken, err := security.VerifyJWT(string(tokenByte), []byte(h.config.JwtSecret))
	if err != nil {
		logger.Error("Error verifying token: %v", err)
		return CustomError[ConfirmEmailResponse]("token is invalid or has expired")
	}

	userIdInterface := verifiedToken.JwtClaims["user_id"]
	userId, err := validateUserIdInterface(userIdInterface)
	if err != nil {
		logger.Error("Error validating user id: %v", err)
		return CustomError[ConfirmEmailResponse](err.Error())
	}

	_, err = h.queries.UpdateUserVerified(c.Request.Context(), database.UpdateUserVerifiedParams{
		ID:       userId,
		Verified: true,
	})
	if err != nil {
		logger.Error("Error updating user: %v", err)
		return GenericError[ConfirmEmailResponse]()
	}

	return &ctx.Response[ConfirmEmailResponse]{
		Response: ConfirmEmailResponse{
			Message: "Email confirmed",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

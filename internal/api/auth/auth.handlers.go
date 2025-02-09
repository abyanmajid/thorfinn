package auth_features

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/matcha/security"
	"github.com/abyanmajid/thorfinn/internal"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type AuthHandlers struct {
	isDev   bool
	config  *internal.EnvConfig
	queries *database.Queries
	mailer  *email.Client
}

func NewHandlers(isDev bool, config *internal.EnvConfig, queries *database.Queries, mailer *email.Client) *AuthHandlers {
	return &AuthHandlers{
		isDev:   isDev,
		config:  config,
		queries: queries,
		mailer:  mailer,
	}
}

func (h *AuthHandlers) Register(c *ctx.Request[RegisterRequest]) *ctx.Response[RegisterResponse] {
	logger.Info("Invoked: Register")

	logger.Debug("Finding user by email")
	_, err := h.queries.FindUserByEmail(c.Request.Context(), c.Body.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("Error finding user by email: %v", err)
		return GenericError[RegisterResponse]()
	}

	if errors.Is(err, sql.ErrNoRows) {
		logger.Debug("Validating register payload")
		err := validateRegisterPayload(c.Body)
		if err != nil {
			logger.Error("Error validating register payload: %v", err)
			return CustomError[RegisterResponse](err.Error())
		}

		logger.Debug("Hashing password")
		passwordHash, err := security.Hash([]byte(c.Body.Password))
		if err != nil {
			logger.Error("Error hashing password: %v", err)
			return GenericError[RegisterResponse]()
		}

		logger.Debug("Creating user")
		user, err := h.queries.CreateUser(c.Request.Context(), database.CreateUserParams{
			ID:           uuid.New().String(),
			Email:        c.Body.Email,
			PasswordHash: string(passwordHash.Hash),
		})
		if err != nil {
			logger.Error("Error creating user: %v", err)
			return GenericError[RegisterResponse]()
		}

		logger.Debug("Creating verification link")
		verificationLink, err := createVerificationLink(VerificationLinkOpts[RegisterRequest]{
			Request: c,
			Config:  h.config,
			UserId:  user.ID,
			Path:    "auth/verify-email",
		})

		if err != nil {
			logger.Error("Error creating verification link: %v", err)
			return GenericError[RegisterResponse]()
		}

		err = h.mailer.SendEmail(h.config.EmailFrom, []string{c.Body.Email}, "Email Verification", "email_verification", map[string]any{
			"VerificationLink": verificationLink,
		})

		if err != nil {
			logger.Error("Error sending email verification email: %v", err)
			return GenericError[RegisterResponse]()
		}
	}

	return &ctx.Response[RegisterResponse]{
		Response: RegisterResponse{
			Message: "If this email is not already in use, you will receive a confirmation email shortly.",
		},
		StatusCode: http.StatusCreated,
		Error:      nil,
	}
}

func (h *AuthHandlers) VerifyEmail(c *ctx.Request[ConfirmEmailRequest]) *ctx.Response[ConfirmEmailResponse] {
	logger.Info("Invoked: VerifyEmail")

	logger.Debug("Decoding, decrypting, and verifying token")
	claims, err := processVerificationToken(c.Body.Token, h.config)
	if err != nil {
		logger.Error("Error processing verification token: %v", err)
		return CustomError[ConfirmEmailResponse](err.Error())
	}

	logger.Debug("Validating user id")
	userIdInterface := claims["user_id"]
	userId, err := validateUserIdInterface(userIdInterface)
	if err != nil {
		logger.Error("Error validating user id: %v", err)
		return CustomError[ConfirmEmailResponse](err.Error())
	}

	logger.Debug("Updating user verification status")
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
			Message: "Email has been verified",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *AuthHandlers) Login(c *ctx.Request[LoginRequest]) *ctx.Response[LoginResponse] {
	logger.Info("Invoked: Login")

	logger.Debug("Finding user by email")
	user, err := h.queries.FindUserByEmail(c.Request.Context(), c.Body.Email)
	if err != nil {
		logger.Error("Error finding user by email: %v", err)
		return GenericError[LoginResponse]()
	}

	if !user.Verified {
		logger.Error("User is not verified")
		return CustomError[LoginResponse]("please verify your email to login")
	}

	logger.Debug("Comparing password with hash")
	err = security.VerifyHash([]byte(user.PasswordHash), []byte(c.Body.Password))
	if err != nil {
		logger.Error("Error verifying password: %v", err)
		return CustomError[LoginResponse]("invalid credentials")
	}

	accessToken, err := h.createAccessToken(&user)
	if err != nil {
		logger.Error("Error creating access token: %v", err)
		return GenericError[LoginResponse]()
	}

	refreshToken, err := h.createRefreshToken(&user)
	if err != nil {
		logger.Error("Error creating refresh token: %v", err)
		return GenericError[LoginResponse]()
	}

	logger.Debug("Setting auth cookies")
	h.setAuthCookies(c, accessToken, refreshToken)

	return &ctx.Response[LoginResponse]{
		Response: LoginResponse{
			Message:       "We have successfully logged you in",
			AccessToken:   accessToken,
			RefreshTokens: refreshToken,
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *AuthHandlers) Logout(c *ctx.Request[LogoutRequest]) *ctx.Response[LogoutResponse] {
	logger.Info("Invoked: Logout")

	logger.Debug("Clearing auth cookies")
	h.clearAuthCookies(c)

	return &ctx.Response[LogoutResponse]{
		Response: LogoutResponse{
			Message: "We have successfully logged you out",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *AuthHandlers) SendEmailVerification(c *ctx.Request[SendVerificationEmailRequest]) *ctx.Response[SendVerificationEmailResponse] {
	logger.Info("Invoked: SendVerificationEmail")

	logger.Debug("Finding user by email")
	user, err := h.queries.FindUserByEmail(c.Request.Context(), c.Body.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("Error finding user by email: %v", err)
		return GenericError[SendVerificationEmailResponse]()
	}

	if !errors.Is(err, sql.ErrNoRows) {
		if user.Verified {
			logger.Error("User is already verified")
			return CustomError[SendVerificationEmailResponse]("user is already verified")
		}

		logger.Debug("Creating verification link")
		verificationLink, err := createVerificationLink(VerificationLinkOpts[SendVerificationEmailRequest]{
			Request: c,
			Config:  h.config,
			UserId:  user.ID,
			Path:    "auth/verify-email",
		})

		if err != nil {
			logger.Error("Error creating verification link: %v", err)
			return GenericError[SendVerificationEmailResponse]()
		}

		err = h.mailer.SendEmail(h.config.EmailFrom, []string{c.Body.Email}, "Email Verification", "email_verification", map[string]any{
			"VerificationLink": verificationLink,
		})

		if err != nil {
			logger.Error("Error sending email verification email: %v", err)
			return GenericError[SendVerificationEmailResponse]()
		}
	}

	return &ctx.Response[SendVerificationEmailResponse]{
		Response: SendVerificationEmailResponse{
			Message: "If this user exists, you will receive a verification email shortly.",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *AuthHandlers) SendPasswordResetLink(c *ctx.Request[SendPasswordResetRequest]) *ctx.Response[SendPasswordResetResponse] {
	logger.Info("Invoked: SendPasswordResetVerification")

	logger.Debug("Finding user by email")
	user, err := h.queries.FindUserByEmail(c.Request.Context(), c.Body.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		logger.Error("Error finding user by email: %v", err)
		return GenericError[SendPasswordResetResponse]()
	}

	if !errors.Is(err, sql.ErrNoRows) {

		logger.Debug("Creating verification link")
		verificationLink, err := createVerificationLink(VerificationLinkOpts[SendPasswordResetRequest]{
			Request: c,
			Config:  h.config,
			UserId:  user.ID,
			Path:    "auth/reset-password",
		})

		if err != nil {
			logger.Error("Error creating verification link: %v", err)
			return GenericError[SendPasswordResetResponse]()
		}

		err = h.mailer.SendEmail(h.config.EmailFrom, []string{c.Body.Email}, "Password Reset", "password_reset_verification", map[string]any{
			"VerificationLink": verificationLink,
		})

		if err != nil {
			logger.Error("Error sending password reset verification email: %v", err)
			return GenericError[SendPasswordResetResponse]()
		}
	}

	return &ctx.Response[SendPasswordResetResponse]{
		Response: SendPasswordResetResponse{
			Message: "If this user exists, you will receive a password reset email shortly.",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *AuthHandlers) ResetPassword(c *ctx.Request[ResetPasswordRequest]) *ctx.Response[ResetPasswordResponse] {
	logger.Info("Invoked: ResetPassword")

	logger.Debug("Decoding, decrypting, and verifying token")
	claims, err := processVerificationToken(c.Body.Token, h.config)
	if err != nil {
		logger.Error("Error processing verification token: %v", err)
		return CustomError[ResetPasswordResponse](err.Error())
	}

	logger.Debug("Validating user id")
	userIdInterface := claims["user_id"]
	userId, err := validateUserIdInterface(userIdInterface)
	if err != nil {
		logger.Error("Error validating user id: %v", err)
		return CustomError[ResetPasswordResponse](err.Error())
	}

	newPasswordHash, err := security.Hash([]byte(c.Body.NewPassword))
	if err != nil {
		logger.Error("Error hashing new password: %v", err)
		return GenericError[ResetPasswordResponse]()
	}

	logger.Debug("Updating user password")
	_, err = h.queries.UpdateUserPassword(c.Request.Context(), database.UpdateUserPasswordParams{
		ID:           userId,
		PasswordHash: newPasswordHash.Hash,
	})
	if err != nil {
		logger.Error("Error updating user password: %v", err)
		return GenericError[ResetPasswordResponse]()
	}

	return &ctx.Response[ResetPasswordResponse]{
		Response: ResetPasswordResponse{
			Message: "Password has been reset",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

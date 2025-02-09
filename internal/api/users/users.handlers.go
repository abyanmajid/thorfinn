package users_features

import (
	"net/http"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/thorfinn/internal"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type UsersHandlers struct {
	isDev   bool
	config  *internal.EnvConfig
	queries *database.Queries
	mailer  *email.Client
}

func NewHandlers(isDev bool, config *internal.EnvConfig, queries *database.Queries, mailer *email.Client) *UsersHandlers {
	return &UsersHandlers{
		isDev:   isDev,
		config:  config,
		queries: queries,
		mailer:  mailer,
	}
}

func (h *UsersHandlers) GetAllUsers(c *ctx.Request[GetAllUsersRequest]) *ctx.Response[GetAllUsersResponse] {
	logger.Info("Invoked: GetAllUsers")

	logger.Debug("Fetching all users")
	users, err := h.queries.ListUsers(c.Request.Context())
	if err != nil {
		logger.Error("Error getting all users: %v", err)
		return internal.GenericError[GetAllUsersResponse]()
	}

	return &ctx.Response[GetAllUsersResponse]{
		Response: GetAllUsersResponse{
			Message: "Successfully fetched all users",
			Users:   users,
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *UsersHandlers) GetUser(c *ctx.Request[GetUserRequest]) *ctx.Response[GetUserResponse] {
	logger.Info("Invoked: GetUser")

	userId := c.GetPathParam("id")

	logger.Debug("Fetching user by id")
	user, err := h.queries.FindUserById(c.Request.Context(), userId)
	if err != nil {
		logger.Error("Error getting user: %v", err)
		return internal.GenericError[GetUserResponse]()
	}

	return &ctx.Response[GetUserResponse]{
		Response: GetUserResponse{
			Message: "Successfully fetched user",
			User:    user,
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

func (h *UsersHandlers) UpdateUser(c *ctx.Request[UpdateUserRequest]) *ctx.Response[UpdateUserResponse] {
	logger.Info("Invoked: UpdateUser")

	userId := c.GetPathParam("id")

	logger.Debug("Fetching existing user details")
	existingUser, err := h.queries.FindUserById(c.Request.Context(), userId)
	if err != nil {
		logger.Error("Error getting user: %v", err)
		return internal.GenericError[UpdateUserResponse]()
	}

	email := existingUser.Email
	if c.Body.Email != nil {
		email = *c.Body.Email
	}

	password := existingUser.PasswordHash
	if c.Body.Password != nil {
		password = *c.Body.Password
	}

	verified := existingUser.Verified
	if c.Body.Verified != nil {
		verified = *c.Body.Verified
	}

	twoFactorEnabled := existingUser.TwoFactorEnabled
	if c.Body.TwoFactorEnabled != nil {
		twoFactorEnabled = *c.Body.TwoFactorEnabled
	}

	logger.Debug("Updating user")
	_, err = h.queries.UpdateUser(c.Request.Context(), database.UpdateUserParams{
		ID:               userId,
		Email:            email,
		PasswordHash:     password,
		Verified:         verified,
		TwoFactorEnabled: twoFactorEnabled,
	})
	if err != nil {
		logger.Error("Error updating user: %v", err)
		return internal.GenericError[UpdateUserResponse]()
	}

	return &ctx.Response[UpdateUserResponse]{
		Response: UpdateUserResponse{
			Message: "Successfully updated user",
		},
		StatusCode: http.StatusOK,
		Error:      nil,
	}
}

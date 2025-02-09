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

	logger.Debug("Getting all users")
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

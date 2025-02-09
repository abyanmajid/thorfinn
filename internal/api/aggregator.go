package api

import (
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/thorfinn/internal"
	auth_features "github.com/abyanmajid/thorfinn/internal/api/auth"
	users_features "github.com/abyanmajid/thorfinn/internal/api/users"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type Resources struct {
	authResources  *auth_features.DerivedAuthResources
	usersResources *users_features.DerivedUsersResources
}

type Handlers struct {
	authHandlers  *auth_features.AuthHandlers
	usersHandlers *users_features.UsersHandlers
}

func aggregateHandlers(isDev bool, config *internal.EnvConfig, queries *database.Queries, mailer *email.Client) *Handlers {
	return &Handlers{
		authHandlers:  auth_features.NewHandlers(isDev, config, queries, mailer),
		usersHandlers: users_features.NewHandlers(isDev, config, queries, mailer),
	}
}

func aggregateResources(handlers *Handlers) (*Resources, error) {
	derivedAuthResources, err := auth_features.Derive(handlers.authHandlers)
	if err != nil {
		return nil, err
	}

	derivedUsersResources, err := users_features.Derive(handlers.usersHandlers)
	if err != nil {
		return nil, err
	}

	return &Resources{
		authResources:  derivedAuthResources,
		usersResources: derivedUsersResources,
	}, nil
}

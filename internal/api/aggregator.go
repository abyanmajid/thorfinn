package api

import (
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/thorfinn/internal"
	auth_features "github.com/abyanmajid/thorfinn/internal/api/auth"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type Resources struct {
	authResources *auth_features.DerivedAuthResources
}

type Handlers struct {
	authHandlers *auth_features.AuthHandlers
}

func aggregateHandlers(isDev bool, config *internal.EnvConfig, queries *database.Queries, mailer *email.Client) *Handlers {
	authHandlers := auth_features.NewHandlers(isDev, config, queries, mailer)

	return &Handlers{
		authHandlers: authHandlers,
	}
}

func aggregateResources(handlers *Handlers) (*Resources, error) {
	derivedAuthResources, err := auth_features.Derive(handlers.authHandlers)
	if err != nil {
		return nil, err
	}

	return &Resources{
		authResources: derivedAuthResources,
	}, nil
}

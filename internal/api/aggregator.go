package api

import (
	"github.com/abyanmajid/thorfinn/internal"
	auth_features "github.com/abyanmajid/thorfinn/internal/api/auth"
)

type Resources struct {
	authResources *auth_features.DerivedAuthResources
}

type Handlers struct {
	authHandlers *auth_features.AuthHandlers
}

func aggregateHandlers(config *internal.EnvConfig) *Handlers {
	authHandlers := auth_features.NewHandlers(config)

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

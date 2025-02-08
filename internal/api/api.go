package api

import (
	"github.com/abyanmajid/thorfinn/internal"
	auth_features "github.com/abyanmajid/thorfinn/internal/api/auth"
)

type ApiResources struct {
	AuthResources *auth_features.DerivedAuthResources
}

func CreateApiResources(config *internal.EnvConfig) (*ApiResources, error) {
	handlers := aggregateHandlers(config)
	resources, err := aggregateResources(handlers)
	if err != nil {
		return nil, err
	}

	return &ApiResources{
		AuthResources: resources.authResources,
	}, nil
}

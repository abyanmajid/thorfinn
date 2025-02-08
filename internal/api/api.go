package api

import (
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/thorfinn/internal"
	auth_features "github.com/abyanmajid/thorfinn/internal/api/auth"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type ApiResources struct {
	AuthResources *auth_features.DerivedAuthResources
}

type Utils struct {
	Config  *internal.EnvConfig
	Queries *database.Queries
	Mailer  *email.Client
}

func CreateApiResources(utils *Utils) (*ApiResources, error) {
	handlers := aggregateHandlers(utils.Config, utils.Queries, utils.Mailer)
	resources, err := aggregateResources(handlers)
	if err != nil {
		return nil, err
	}

	return &ApiResources{
		AuthResources: resources.authResources,
	}, nil
}

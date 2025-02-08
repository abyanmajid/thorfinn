package auth_features

import "github.com/abyanmajid/matcha/openapi"

type DerivedAuthResources struct {
	Register    *openapi.Resource
	VerifyEmail *openapi.Resource
}

func Derive(handlers *AuthHandlers) (*DerivedAuthResources, error) {
	authResources := NewAuthResources(handlers)
	registerResource, err := authResources.RegisterResource()
	if err != nil {
		return nil, err
	}

	confirmEmailResource, err := authResources.ConfirmEmailResource()
	if err != nil {
		return nil, err
	}

	return &DerivedAuthResources{
		Register:    registerResource,
		VerifyEmail: confirmEmailResource,
	}, nil
}

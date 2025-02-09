package users_features

import "github.com/abyanmajid/matcha/openapi"

type DerivedUsersResources struct {
	GetAllUsers *openapi.Resource
	GetUser     *openapi.Resource
	UpdateUser  *openapi.Resource
}

func Derive(handlers *UsersHandlers) (*DerivedUsersResources, error) {
	userResources := NewUsersResources(handlers)
	getAllUsersResource, err := userResources.GetAllUsersResource()
	if err != nil {
		return nil, err
	}

	getUserResource, err := userResources.GetUserResource()
	if err != nil {
		return nil, err
	}

	updateUserResource, err := userResources.UpdateUserResource()
	if err != nil {
		return nil, err
	}

	return &DerivedUsersResources{
		GetAllUsers: getAllUsersResource,
		GetUser:     getUserResource,
		UpdateUser:  updateUserResource,
	}, nil
}

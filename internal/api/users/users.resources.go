package users_features

import (
	"net/http"

	"github.com/abyanmajid/matcha/openapi"
)

type UsersResources struct {
	handlers *UsersHandlers
}

func NewUsersResources(handlers *UsersHandlers) *UsersResources {
	return &UsersResources{
		handlers: handlers,
	}
}

func (r *UsersResources) GetAllUsersResource() (*openapi.Resource, error) {
	requestSchema, err := openapi.NewSchema(GetAllUsersRequest{})
	if err != nil {
		return nil, err
	}

	responseSchema, err := openapi.NewSchema(GetAllUsersResponse{})
	if err != nil {
		return nil, err
	}

	doc := openapi.ResourceDoc{
		Summary:     "Get all users",
		Description: "Get all users",
		Schema: openapi.Schema{
			RequestBody: openapi.RequestBody{
				Content: openapi.Json(requestSchema),
			},
			Responses: map[int]openapi.Response{
				http.StatusOK: {
					Description: "Successfully fetched all users",
					Content:     openapi.Json(responseSchema),
				},
			},
		},
	}

	resource := openapi.NewResource("GetAllUsers", doc, r.handlers.GetAllUsers)

	return &resource, nil
}

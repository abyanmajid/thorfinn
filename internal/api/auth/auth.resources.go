package auth_features

import (
	"net/http"

	"github.com/abyanmajid/matcha/openapi"
)

type AuthResources struct {
	handlers *AuthHandlers
}

func NewAuthResources(handlers *AuthHandlers) *AuthResources {
	return &AuthResources{
		handlers: handlers,
	}
}

func (r *AuthResources) RegisterResource() (*openapi.Resource, error) {
	requestSchema, err := openapi.NewSchema(RegisterRequest{})
	if err != nil {
		return nil, err
	}

	responseSchema, err := openapi.NewSchema(RegisterResponse{})
	if err != nil {
		return nil, err
	}

	doc := openapi.ResourceDoc{
		Summary:     "Register a new user",
		Description: "Check if a user exists, if not, validate the inputs, and send a confirmation email to the user",
		Schema: openapi.Schema{
			RequestBody: openapi.RequestBody{
				Content: openapi.Json(requestSchema),
			},
			Responses: map[int]openapi.Response{
				http.StatusOK: {
					Description: "Please check your email for a verification link",
					Content:     openapi.Json(responseSchema),
				},
				http.StatusUnauthorized: {
					Description: "Invalid credentials.",
					Content:     openapi.Json(openapi.SimpleErrorSchema()),
				},
			},
		},
	}

	resource := openapi.NewResource("Register", doc, r.handlers.Register)

	return &resource, nil
}

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
			},
		},
	}

	resource := openapi.NewResource("Register", doc, r.handlers.Register)

	return &resource, nil
}

func (r *AuthResources) ConfirmEmailResource() (*openapi.Resource, error) {
	requestSchema, err := openapi.NewSchema(ConfirmEmailRequest{})
	if err != nil {
		return nil, err
	}

	responseSchema, err := openapi.NewSchema(ConfirmEmailResponse{})
	if err != nil {
		return nil, err
	}

	doc := openapi.ResourceDoc{
		Summary:     "Verify email",
		Description: "Verify email",
		Schema: openapi.Schema{
			RequestBody: openapi.RequestBody{Content: openapi.Json(requestSchema)},
			Responses: map[int]openapi.Response{
				http.StatusOK: {
					Description: "Email confirmed",
					Content:     openapi.Json(responseSchema),
				},
			},
		},
	}

	resource := openapi.NewResource("VerifyEmail", doc, r.handlers.VerifyEmail)

	return &resource, nil
}

func (r *AuthResources) LoginResource() (*openapi.Resource, error) {
	requestSchema, err := openapi.NewSchema(LoginRequest{})
	if err != nil {
		return nil, err
	}

	responseSchema, err := openapi.NewSchema(LoginResponse{})
	if err != nil {
		return nil, err
	}

	doc := openapi.ResourceDoc{
		Summary:     "Login",
		Description: "Login",
		Schema: openapi.Schema{
			RequestBody: openapi.RequestBody{Content: openapi.Json(requestSchema)},
			Responses: map[int]openapi.Response{
				http.StatusOK: {
					Description: "Login successful",
					Content:     openapi.Json(responseSchema),
				},
			},
		},
	}

	resource := openapi.NewResource("Login", doc, r.handlers.Login)

	return &resource, nil
}

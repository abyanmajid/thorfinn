package users_features

import (
	"github.com/abyanmajid/thorfinn/internal/database"
)

type GetAllUsersRequest struct{}

type GetAllUsersResponse struct {
	Message string                  `json:"message"`
	Users   []database.ListUsersRow `json:"users"`
}

type GetUserRequest struct{}

type GetUserResponse struct {
	Message string                `json:"message"`
	User    database.ThorfinnUser `json:"user"`
}

type UpdateUserRequest struct {
	Email            *string `json:"email,omitempty"`
	Password         *string `json:"password,omitempty"`
	Verified         *bool   `json:"verified,omitempty"`
	TwoFactorEnabled *bool   `json:"two_factor_enabled,omitempty"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
}

type DeleteUserRequest struct{}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

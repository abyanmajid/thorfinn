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
	ID string `json:"id"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

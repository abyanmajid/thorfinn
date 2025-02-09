package users_features

import (
	"github.com/abyanmajid/thorfinn/internal/database"
)

type GetAllUsersRequest struct{}

type GetAllUsersResponse struct {
	Message string                  `json:"message"`
	Users   []database.ListUsersRow `json:"users"`
}

type GetUserRequest struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	Message string `json:"message"`
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

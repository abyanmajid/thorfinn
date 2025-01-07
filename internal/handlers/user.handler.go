package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lucsky/cuid"

	"github.com/clyde-sh/orion/internal/database"
	"github.com/clyde-sh/orion/internal/security"
	"github.com/clyde-sh/orion/internal/utils"
	shared "github.com/clyde-sh/orion/shared/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

func (h *HandlersCtx) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.UserRegisterRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, ErrRateLimitedError, http.StatusTooManyRequests)
		return
	}

	var body shared.UserCreateUserRequestDto
	err := utils.ReadJSON(w, r, &body)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.IsValidEmail(body.Email)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.IsValidPassword(body.Password)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = utils.IsValidName(body.Name)
	if err != nil {
		utils.WriteErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	passwordHash, err := security.HashValue(body.Password)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	accountRecoveryCode, err := security.GenerateSecureCode()
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	accountRecoveryCodeHash, err := security.HashValue(accountRecoveryCode)
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	id := cuid.New()

	slog.Debug(fmt.Sprintf("Inserting a new user of id %s...", id))

	user, err := h.queries.InsertUser(r.Context(), database.InsertUserParams{
		ID:           id,
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: passwordHash,
		Role:         string(RoleUser),
		RecoveryCode: accountRecoveryCodeHash,
	})
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, shared.UserCreateUserResponseDto{
		User: &shared.User{
			Id:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Role:            user.Role,
			IsEmailVerified: user.IsEmailVerified.Bool,
			CreatedAt:       timestamppb.New(user.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(user.UpdatedAt.Time),
		},
	}, http.StatusCreated)
}

func (h *HandlersCtx) HandleAllGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.queries.ListUsers(r.Context())
	if err != nil {
		utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	var protoUsers []*shared.User
	for _, user := range users {
		pu := &shared.User{
			Id:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Role:            user.Role,
			IsEmailVerified: user.IsEmailVerified.Bool,
			Is_2FaEnabled:   user.Is2faEnabled.Bool,
			CreatedAt:       timestamppb.New(user.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(user.UpdatedAt.Time),
		}
		protoUsers = append(protoUsers, pu)
	}

	utils.WriteJSON(w, &shared.UserGetAllUsersResponseDto{
		Users: protoUsers,
	}, http.StatusOK)
}

func (h *HandlersCtx) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.queries.FindUserById(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.WriteErrorJSON(w, ErrNotFoundError, http.StatusNotFound)
		} else {
			utils.WriteErrorJSON(w, ErrInternalServerError, http.StatusInternalServerError)
		}
		return
	}

	utils.WriteJSON(w, shared.UserCreateUserResponseDto{
		User: &shared.User{
			Id:              user.ID,
			Name:            user.Name,
			Email:           user.Email,
			Role:            user.Role,
			IsEmailVerified: user.IsEmailVerified.Bool,
			CreatedAt:       timestamppb.New(user.CreatedAt.Time),
			UpdatedAt:       timestamppb.New(user.UpdatedAt.Time),
		},
	}, http.StatusCreated)
}

func (h *HandlersCtx) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {

}

func (h *HandlersCtx) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {

}

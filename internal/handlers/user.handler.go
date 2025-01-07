package handlers

import (
	"errors"
	"net/http"

	"github.com/lucsky/cuid"

	"github.com/clyde-sh/orion/internal/database"
	"github.com/clyde-sh/orion/internal/security"
	"github.com/clyde-sh/orion/internal/utils"
	shared "github.com/clyde-sh/orion/shared/go"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type UserHandlers struct {
	queries   *database.Queries
	ratelimit *security.AppRateLimiters
}

func NewUserHandlers(queries *database.Queries, appRateLimiters *security.AppRateLimiters) UserHandlers {
	return UserHandlers{
		queries:   queries,
		ratelimit: appRateLimiters,
	}
}

func (h *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetClientIP(r)

	if ip != "" && !h.ratelimit.AuthRegisterRateLimit.Consume(ip) {
		utils.WriteErrorJSON(w, security.ErrRateLimited, http.StatusTooManyRequests)
		return
	}

	var body shared.AuthRegisterRequestDto
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
		utils.WriteErrorJSON(w, security.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	accountRecoveryCode, err := security.GenerateSecureCode()
	if err != nil {
		utils.WriteErrorJSON(w, security.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	accountRecoveryCodeHash, err := security.HashValue(accountRecoveryCode)
	if err != nil {
		utils.WriteErrorJSON(w, security.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	err = h.queries.InsertUser(r.Context(), database.InsertUserParams{
		ID:           cuid.New(),
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: passwordHash,
		Role:         string(RoleUser),
		RecoveryCode: accountRecoveryCodeHash,
	})
	if err != nil {
		utils.WriteErrorJSON(w, errors.New("an error occurred while issuing your account"), http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, shared.AuthRegisterResponseDto{
		Status:  http.StatusCreated,
		Message: "Successfully created a user.",
	}, http.StatusCreated)
}

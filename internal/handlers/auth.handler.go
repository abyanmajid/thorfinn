package handlers

import (
	"net/http"

	"github.com/clyde-sh/novus/internal/database"
	"github.com/clyde-sh/novus/internal/utils"
)

type AuthHandlers struct {
	queries *database.Queries
}

func NewAuthHandlers(queries *database.Queries) AuthHandlers {
	return AuthHandlers{
		queries: queries,
	}
}

func (h *AuthHandlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.GetClientIP(r)
}

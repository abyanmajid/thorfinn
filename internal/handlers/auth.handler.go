package handlers

import (
	"net/http"

	"github.com/clyde-sh/novus/internal/database"
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

}

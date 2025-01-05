package handlers

import "net/http"

type AuthHandlers struct {
	DbUrl string
}

func NewAuthHandlers(dbUrl string) AuthHandlers {
	return AuthHandlers{
		DbUrl: dbUrl,
	}
}

func (h *AuthHandlers) RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

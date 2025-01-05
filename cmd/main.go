package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/clyde-sh/novus/internal/handlers"
	"github.com/clyde-sh/novus/internal/utils"
	"github.com/go-chi/chi/v5"
)

type NovusAPI struct {
	port  int
	dbUrl string
}

func (cfg *NovusAPI) New() chi.Router {
	authHandlers := handlers.NewAuthHandlers(cfg.dbUrl)

	api := chi.NewRouter()
	api.Get("/auth/register", authHandlers.RegisterHandler)

	return api
}

func main() {
	slog.Debug("Loading environment variables and registering routes...")

	env, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	novus := &NovusAPI{
		port:  env.PORT,
		dbUrl: env.DATABASE_URL,
	}

	r := chi.NewRouter()

	r.Mount("/api", novus.New())

	slog.Info("Clyde Noxus API is now running live on port 8080...")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", env.PORT), r); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}

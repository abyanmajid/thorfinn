package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/clyde-sh/novus/internal/database"
	"github.com/clyde-sh/novus/internal/handlers"
	"github.com/clyde-sh/novus/internal/utils"
	"github.com/go-chi/chi/v5"
)

type NovusAPI struct {
	port    int
	queries *database.Queries
}

func (cfg *NovusAPI) NewAPI() chi.Router {
	userHandlers := handlers.NewUserHandlers(cfg.queries)

	api := chi.NewRouter()
	api.Post("/users", userHandlers.HandleCreateUser)

	return api
}

func main() {
	env, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	queries, err := utils.CreateQueryClient()
	if err != nil {
		log.Fatal(err)
	}

	novus := &NovusAPI{
		port:    env.PORT,
		queries: queries,
	}

	go utils.ScheduleDailyDatabaseCleanUp(queries)

	r := chi.NewRouter()
	r.Mount("/api", novus.NewAPI())

	slog.Info(fmt.Sprintf("Clyde Noxus API is now running live on port %d...", novus.port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", novus.port), r); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}

package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/clyde-sh/novus/internal/database"
	"github.com/clyde-sh/novus/internal/handlers"
	"github.com/clyde-sh/novus/internal/security"
	"github.com/clyde-sh/novus/internal/utils"
	"github.com/go-chi/chi/v5"
)

type NovusAPI struct {
	env             *utils.Env
	appRateLimiters security.AppRateLimiters
	queries         *database.Queries
}

func (cfg *NovusAPI) NewAPI() chi.Router {
	userHandlers := handlers.NewUserHandlers(cfg.queries, &cfg.appRateLimiters)

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

	appRateLimiters := security.InitializeAppRateLimiters()

	novus := &NovusAPI{
		env:             env,
		queries:         queries,
		appRateLimiters: appRateLimiters,
	}

	go utils.ScheduleDailyDatabaseCleanUp(queries)

	r := chi.NewRouter()
	r.Mount("/api", novus.NewAPI())

	slog.Info(fmt.Sprintf("Clyde Noxus API is now running live on port %d...", novus.env.PORT))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", novus.env.PORT), r); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}

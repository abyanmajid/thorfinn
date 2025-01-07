package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/clyde-sh/orion/internal/database"
	"github.com/clyde-sh/orion/internal/handlers"
	"github.com/clyde-sh/orion/internal/security"
	"github.com/clyde-sh/orion/internal/utils"
	"github.com/go-chi/chi/v5"
)

type OrionAPI struct {
	env             *utils.Env
	appRateLimiters security.AppRateLimiters
	queries         *database.Queries
}

func (cfg *OrionAPI) NewAPI() chi.Router {
	handlersCtx := handlers.New(cfg.queries, &cfg.appRateLimiters)

	api := chi.NewRouter()
	api.Post("/users", handlersCtx.HandleCreateUser)
	api.Get("/users", handlersCtx.HandleAllGetUsers)
	api.Get("/users/:id", handlersCtx.HandleGetUserById)
	api.Put("/users/:id", handlersCtx.HandleUpdateUserById)
	api.Delete("users/:id", handlersCtx.HandleDeleteUserById)

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

	orion := &OrionAPI{
		env:             env,
		queries:         queries,
		appRateLimiters: appRateLimiters,
	}

	go utils.ScheduleDailyDatabaseCleanUp(queries)

	r := chi.NewRouter()
	r.Mount("/api", orion.NewAPI())

	slog.Info(fmt.Sprintf("Clyde Noxus API is now running live on port %d...", orion.env.PORT))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", orion.env.PORT), r); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}

package main

import (
	"github.com/abyanmajid/matcha"
	"github.com/abyanmajid/matcha/email"
	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/matcha/openapi"
	"github.com/abyanmajid/matcha/reference"
	"github.com/abyanmajid/thorfinn/internal"
	"github.com/abyanmajid/thorfinn/internal/api"
)

func main() {
	config := internal.ConfigureEnv()

	queries, err := internal.CreateQueryClient(config.DatabaseUrl)
	if err != nil {
		logger.Fatal("Failed to create query client: %v", err)
	}

	mailer := email.NewClient(email.Config{
		Host:     config.SmtpHost,
		Port:     config.SmtpPort,
		Username: config.SmtpUser,
		Password: config.SmtpPassword,
	}, "templates")

	resources, err := api.CreateApiResources(&api.Utils{
		Config:  config,
		Queries: queries,
		Mailer:  mailer,
	})
	if err != nil {
		logger.Fatal("Failed to create resources: %v", err)
	}

	app := matcha.New()

	app.Documentation("/docs", openapi.Meta{
		OpenAPI:        "3.0.0",
		PackageName:    "Thorfinn API",
		PackageVersion: "0.1.0",
	})

	app.Post("/auth/register", resources.AuthResources.Register)

	app.Reference("/reference", &reference.Options{
		Source: "/docs",
	})

	app.Serve(":8080")
}

package utils

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	ENVIRONMENT  string
	PORT         int
	DATABASE_URL string
}

func LoadEnv() (*Env, error) {
	slog.Debug("Loading environment variables...")

	env := "production"
	for _, arg := range os.Args {
		if arg == "--env=development" || arg == "-e=development" {
			env = "development"
			break
		} else if arg == "--env=production" || arg == "-e=production" {
			env = "production"
			break
		}
	}

	slog.Debug(fmt.Sprintf("Configured app to run in %s mode.", env))

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	portStr := os.Getenv("PORT")

	port := 8080
	if portStr != "" {
		parsedPort, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("PORT is invalid: %v", err)
		}
		port = parsedPort
	}

	slog.Debug(fmt.Sprintf("Configured app to run on port %d.", port))

	dbUrl, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		return nil, fmt.Errorf("DATABASE_URL is required, but is unset")
	}

	return &Env{
		PORT:         port,
		DATABASE_URL: dbUrl,
	}, nil
}

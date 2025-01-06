package utils

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	ENVIRONMENT       string
	DATABASE_URL      string
	PORT              int
	JWT_SECRET        string
	ENCRYPTION_SECRET string
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
		} else if arg == "--env=development-serverless" || arg == "-e=development-serverless" {
			env = "development-serverless"
			break
		} else if arg == "--env=production-serverless" || arg == "-e=production-serverless" {
			env = "production-serverless"
			break
		}
	}

	if env != "development" && env != "production" && env != "development-serverless" && env != "production-serverless" {
		log.Fatalf("Invalid environment specified.")
	}

	slog.Debug(fmt.Sprintf("Configured app to run in %s mode.", env))

	if strings.HasPrefix(env, "development") {
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

	jwtSecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return nil, fmt.Errorf("JWT_SECRET is required, but is unset")
	}

	encryptionSecret, exists := os.LookupEnv("ENCRYPTION_SECRET")
	if !exists {
		return nil, fmt.Errorf("ENCRYPTION_SECRET is required, but is unset")
	}

	return &Env{
		ENVIRONMENT:       env,
		PORT:              port,
		DATABASE_URL:      dbUrl,
		JWT_SECRET:        jwtSecret,
		ENCRYPTION_SECRET: encryptionSecret,
	}, nil
}

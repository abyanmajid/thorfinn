package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	PORT         int
	DATABASE_URL string
}

func LoadEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	dbUrl, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		return nil, fmt.Errorf("DATABASE_URL is required, but is unset")
	}

	return &Env{
		PORT:         port,
		DATABASE_URL: dbUrl,
	}, nil
}

package internal

import (
	"context"

	"github.com/abyanmajid/matcha/logger"
	"github.com/abyanmajid/thorfinn/internal/database"
	"github.com/jackc/pgx/v5"
)

func CreateQueryClient(connStr string) (*database.Queries, error) {
	logger.Debug("Establishing database connection...")

	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	logger.Info("Successfully established database connection.")

	queries := database.New(db)

	return queries, nil
}

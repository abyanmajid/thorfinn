package utils

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/clyde-sh/orion/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func CleanUpDatabase(queries *database.Queries) error {
	now := &pgtype.Timestamptz{Time: time.Now(), Valid: true}

	slog.Info("Cleaning up expired emails verification requests...")
	err := queries.DeleteExpiredEmailVerificationRequests(context.Background(), *now)
	if err != nil {
		return err
	}

	slog.Info("Cleaning up expired password verification requests...")
	err = queries.DeleteExpiredPasswordResetRequests(context.Background(), *now)
	if err != nil {
		return err
	}

	return nil
}

func CreateQueryClient() (*database.Queries, error) {
	slog.Debug("Establishing database connection...")

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	slog.Info("Successfully established database connection.")

	defer conn.Close(context.Background())

	queries := database.New(conn)

	return queries, nil
}

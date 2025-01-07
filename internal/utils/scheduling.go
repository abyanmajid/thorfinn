package utils

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/clyde-sh/orion/internal/database"
)

func ScheduleDailyDatabaseCleanUp(queries *database.Queries) {
	for range time.Tick(24 * time.Hour) {
		err := CleanUpDatabase(queries)
		if err != nil {
			slog.Error(fmt.Sprintf("Failed to perform daily scheduled database cleanup: %v", err))
		}
	}
}

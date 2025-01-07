package handlers

import (
	"github.com/clyde-sh/orion/internal/database"
	"github.com/clyde-sh/orion/internal/security"
)

type HandlersCtx struct {
	queries   *database.Queries
	ratelimit *security.AppRateLimiters
}

func New(queries *database.Queries, ratelimit *security.AppRateLimiters) HandlersCtx {
	return HandlersCtx{
		queries:   queries,
		ratelimit: ratelimit,
	}
}

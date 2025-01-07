package handlers

import (
	"errors"

	"github.com/clyde-sh/orion/internal/database"
	"github.com/clyde-sh/orion/internal/security"
)

var ErrRateLimitedError = errors.New("too many requests")
var ErrInternalServerError = errors.New("something went wrong")
var ErrNotFoundError = errors.New("not found")

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

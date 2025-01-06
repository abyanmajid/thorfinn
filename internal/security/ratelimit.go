package security

import (
	"errors"
	"time"
)

var ErrRateLimited = errors.New("too many requests")
var ErrInternalServerError = errors.New("something went wrong")

type AppRateLimiters struct {
	AuthRegisterRateLimit TokenBucketRateLimit
}

func InitializeAppRateLimiters() AppRateLimiters {
	return AppRateLimiters{
		AuthRegisterRateLimit: NewTokenBucketRateLimit(5, 10*time.Second),
	}
}

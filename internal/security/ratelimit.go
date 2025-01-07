package security

import (
	"time"
)

type AppRateLimiters struct {
	UserRegisterRateLimit    TokenBucketRateLimit
	UserGetUserByIdRateLimit TokenBucketRateLimit
	UserGetUsersRateLimit    TokenBucketRateLimit
}

func InitializeAppRateLimiters() AppRateLimiters {
	return AppRateLimiters{
		UserRegisterRateLimit:    NewTokenBucketRateLimit(5, 10*time.Second),
		UserGetUserByIdRateLimit: NewTokenBucketRateLimit(50, 1*time.Minute),
		UserGetUsersRateLimit:    NewTokenBucketRateLimit(50, 1*time.Minute),
	}
}

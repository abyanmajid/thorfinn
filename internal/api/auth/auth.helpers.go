package auth_features

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/security"
	"github.com/abyanmajid/thorfinn/internal"
	"github.com/abyanmajid/thorfinn/internal/database"
)

type VerificationLinkOpts[T any] struct {
	Request *ctx.Request[T]
	Config  *internal.EnvConfig
	UserId  string
	Path    string
}

func createVerificationLink[T any](opts VerificationLinkOpts[T]) (string, error) {
	token := security.NewJWT(security.JwtClaims{
		"user_id": opts.UserId,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * 10).Unix(),
	})

	signedToken, err := token.Sign([]byte(opts.Config.JwtSecret))
	if err != nil {
		return "", err
	}

	encryptedSignedToken, err := security.Encrypt([]byte(signedToken), []byte(opts.Config.EncryptionSecret), []byte(opts.Config.EncryptionIv))
	if err != nil {
		return "", err
	}

	tokenUrlSafe := security.EncodeBase64(encryptedSignedToken)

	verificationLink := fmt.Sprintf("%s/%s?token=%s", opts.Config.FrontendUrl, opts.Path, tokenUrlSafe)

	return verificationLink, nil
}

const otpCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateOtp(length int) (string, error) {
	bytes := make([]byte, length)
	for i := range bytes {
		randomByte, err := rand.Int(rand.Reader, big.NewInt(int64(len(otpCharset))))
		if err != nil {
			return "", err
		}
		bytes[i] = otpCharset[randomByte.Int64()]
	}
	return string(bytes), nil
}

func processVerificationToken(token string, config *internal.EnvConfig) (security.JwtClaims, error) {
	if token == "" {
		return nil, errors.New("token not found")
	}

	encryptedToken, err := security.DecodeBase64(token)
	if err != nil {
		return nil, errors.New("an error occurred while processing your request")
	}

	tokenByte, err := security.Decrypt([]byte(encryptedToken), []byte(config.EncryptionSecret))
	if err != nil {
		return nil, errors.New("an error occurred while processing your request")
	}

	verifiedToken, err := security.VerifyJWT(string(tokenByte), []byte(config.JwtSecret))
	if err != nil {
		return nil, errors.New("token is invalid or has expired")
	}

	return verifiedToken.JwtClaims, nil
}

func (h *AuthHandlers) setAuthCookies(c *ctx.Request[LoginRequest], accessToken string, refreshToken string) {
	c.Cookies.SetCookie("access_token", accessToken, &ctx.CookieOptions{
		Path:     "/",
		Domain:   fmt.Sprintf(".%s", h.config.RootDomain),
		MaxAge:   15 * 60,
		Secure:   !h.isDev,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(15 * time.Minute),
	})

	c.Cookies.SetCookie("refresh_token", refreshToken, &ctx.CookieOptions{
		Path:     "/",
		Domain:   fmt.Sprintf(".%s", h.config.RootDomain),
		MaxAge:   30 * 24 * 60 * 60,
		Secure:   !h.isDev,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})
}

func (h *AuthHandlers) clearAuthCookies(c *ctx.Request[LogoutRequest]) {
	c.Cookies.SetCookie("access_token", "", &ctx.CookieOptions{
		Path:     "/",
		Domain:   fmt.Sprintf(".%s", h.config.RootDomain),
		MaxAge:   -1,
		Secure:   !h.isDev,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Unix(0, 0),
	})

	c.Cookies.SetCookie("refresh_token", "", &ctx.CookieOptions{
		Path:     "/",
		Domain:   fmt.Sprintf(".%s", h.config.RootDomain),
		MaxAge:   -1,
		Secure:   !h.isDev,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Unix(0, 0),
	})
}

func (h *AuthHandlers) createAccessToken(user *database.ThorfinnUser) (string, error) {
	claims := security.JwtClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}

	unsignedAccessToken := security.NewJWT(claims)
	signedAccessToken, err := unsignedAccessToken.Sign([]byte(h.config.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("error signing access token: %v", err)
	}

	encryptedSignedAccessToken, err := security.Encrypt([]byte(signedAccessToken), []byte(h.config.EncryptionSecret), []byte(h.config.EncryptionIv))
	if err != nil {
		return "", fmt.Errorf("error encrypting signed access token: %v", err)
	}

	return security.EncodeBase64(encryptedSignedAccessToken), nil
}

func (h *AuthHandlers) createRefreshToken(user *database.ThorfinnUser) (string, error) {
	claims := security.JwtClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	unsignedRefreshToken := security.NewJWT(claims)
	signedRefreshToken, err := unsignedRefreshToken.Sign([]byte(h.config.JwtSecret))
	if err != nil {
		return "", fmt.Errorf("error signing refresh token: %v", err)
	}

	encryptedSignedRefreshToken, err := security.Encrypt([]byte(signedRefreshToken), []byte(h.config.EncryptionSecret), []byte(h.config.EncryptionIv))
	if err != nil {
		return "", fmt.Errorf("error encrypting signed refresh token: %v", err)
	}

	return security.EncodeBase64(encryptedSignedRefreshToken), nil
}

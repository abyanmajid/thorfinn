package auth_features

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abyanmajid/matcha/ctx"
	"github.com/abyanmajid/matcha/security"
	"github.com/abyanmajid/thorfinn/internal/database"
)

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

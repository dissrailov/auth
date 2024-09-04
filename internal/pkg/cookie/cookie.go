package cookie

import (
	"auth/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetAuthCookies(c *gin.Context, refreshToken, accessToken string, cfg config.JWT) {
	setRefreshCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(cfg.RefreshTokenTTL),
		Path:     "/api/auth",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, &setRefreshCookie)
	setAccessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(cfg.AccessTokenTTL),
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, &setAccessTokenCookie)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

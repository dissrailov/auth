package handlers

import (
	"auth/internal/pkg/cookie"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *HandlerApp) Hello(c *gin.Context) {
	fmt.Println("Hello World")
	c.JSON(200, gin.H{"DAMIR": "LEZHAT'"})
	return
}
func (h *HandlerApp) GetAccessToken(c *gin.Context) {
	guidStr := c.Query("guid")
	if guidStr == "" {
		c.JSON(400, gin.H{"error": "GUID is empty"})
		return
	}

	guid, err := uuid.Parse(guidStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid GUID"})
		return
	}
	tokenPair, err := h.service.ForgeAuthPair(guid.String(), c.ClientIP(), *h.cfg)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cookie.SetAuthCookies(c, tokenPair.RefreshToken, tokenPair.AccessToken, h.cfg.JWT)

	c.JSON(200, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}

func (h *HandlerApp) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("refresh_token")
	if refreshToken == "" {
		c.JSON(400, gin.H{"error": "Refresh token is empty"})
		return
	}
	guid := c.GetHeader("guid")
	if guid == "" {
		c.JSON(400, gin.H{"error": "GUID is empty"})
		return
	}
	_, valid, err := h.service.VerifyRefreshToken(refreshToken, guid, c.ClientIP())
	if err != nil || !valid {
		c.JSON(400, gin.H{"error": "Refresh token is invalid or expired"})
		return
	}
	err = h.service.DeleteRefreshToken(guid)
	if err != nil {
		c.JSON(500, gin.H{"error": "cannot delete refresh token"})
	}
	newGUID := uuid.New().String()
	tokenPair, err := h.service.ForgeAuthPair(newGUID, c.ClientIP(), *h.cfg)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	cookie.SetAuthCookies(c, tokenPair.RefreshToken, tokenPair.AccessToken, h.cfg.JWT)
	c.JSON(200, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}

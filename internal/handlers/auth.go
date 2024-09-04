package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

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
	log.Println(c.ClientIP())
	tokenPair, err := h.service.ForgeAuthPair(guid.String(), c.ClientIP(), *h.cfg)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}

func (h *HandlerApp) RefreshToken(c *gin.Context) {
}

package handlers

import "github.com/gin-gonic/gin"

func (h *HandlerApp) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/token")
		auth.POST("/refresh")
	}
	return router
}

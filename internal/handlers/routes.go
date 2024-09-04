package handlers

import "github.com/gin-gonic/gin"

func (h *HandlerApp) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/token", h.GetAuthToken)
	}
	return router
}

package handlers

import "github.com/gin-gonic/gin"

func (h *HandlerApp) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("", h.Hello)
		auth.GET("/access", h.GetAccessToken)
		auth.POST("refresh", h.RefreshToken)
	}
	return router
}

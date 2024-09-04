package handlers

import (
	"auth/internal/config"
	"auth/internal/service"
)

type HandlerApp struct {
	service service.ServiceI
	cfg     *config.Config
}

func NewHandler(s service.ServiceI, c *config.Config) *HandlerApp {
	return &HandlerApp{
		s,
		c,
	}
}

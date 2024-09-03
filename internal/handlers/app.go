package handlers

import "auth/internal/service"

type HandlerApp struct {
	service service.ServiceI
}

func NewHandler(s service.ServiceI) *HandlerApp {
	return &HandlerApp{
		s,
	}
}

package service

import (
	"auth/internal/config"
	"auth/internal/models"
	"auth/internal/repo"
)

type service struct {
	repo repo.RepoI
}
type ServiceI interface {
	Authenticate
}

type Authenticate interface {
	ForgeAuthPair(guid string, ip string, cfg config.Config) (*models.TokenPair, error)
	VerifyRefreshToken(refreshToken, guid, currentIP string) (string, bool, error)
	DeleteRefreshToken(guid string) error
}

func NewService(repo repo.RepoI) ServiceI {
	return &service{repo}
}

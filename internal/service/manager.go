package service

import "auth/internal/repo"

type service struct {
	repo repo.RepoI
}
type ServiceI interface {
	Authenticate
}

type Authenticate interface {
}

func NewService(repo repo.RepoI) ServiceI {
	return &service{repo}
}

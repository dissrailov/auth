package repo

import (
	"auth/internal/repo/postgres"
	"github.com/jmoiron/sqlx"
)

type RepoI interface {
	AuthenticationI
}
type AuthenticationI interface {
	InsertRefreshToken(guid, hashToken, ip string) error
	GetRefreshToken(guid string) (string, string, error)
	DeleteRefreshToken(guid string) error
}

func NewRepository(db *sqlx.DB) RepoI {
	return &postgres.Storage{DB: db}
}

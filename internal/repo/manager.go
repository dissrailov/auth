package repo

import (
	"auth/internal/repo/postgres"
	"github.com/jmoiron/sqlx"
)

type RepoI interface {
	AuthenticationI
}
type AuthenticationI interface {
	InsertRefreshToken(guid, hashToken string) error
}

func NewRepository(db *sqlx.DB) RepoI {
	return &postgres.Storage{DB: db}
}

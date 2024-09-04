package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Email string    `json:"email" db:"email"`
}

type TokenPair struct {
	AccessToken   string     `json:"access_token"`
	AccessObject  *jwt.Token `json:"-"`
	RefreshToken  string     `json:"refresh_token"`
	RefreshObject *jwt.Token `json:"-"`
}

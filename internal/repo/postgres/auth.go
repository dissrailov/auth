package postgres

import (
	"fmt"
	"time"
)

func (s *Storage) InsertRefreshToken(guid, hashToken, ip string) error {
	op := "postgres.InsertRefreshToken"
	query := `INSERT INTO refresh_tokens (guid, hashed_token, created_at, ip) VALUES ($1, $2, $3, $4)`
	_, err := s.DB.Exec(query, guid, hashToken, time.Now(), ip)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Storage) GetRefreshToken(guid string) (string, string, error) {
	op := "postgres.GetRefreshToken"
	var ip string
	var hashedToken string
	query := `SELECT ip,hashed_token FROM refresh_tokens WHERE guid = $1`
	err := s.DB.QueryRow(query, guid).Scan(&ip, &hashedToken)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	return hashedToken, ip, nil
}

func (s *Storage) DeleteRefreshToken(guid string) error {
	op := "postgres.DeleteRefreshToken"
	query := `DELETE FROM refresh_tokens WHERE guid = $1`
	_, err := s.DB.Exec(query, guid)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

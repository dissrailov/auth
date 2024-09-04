package postgres

import "time"

func (s *Storage) InsertRefreshToken(guid, hashToken string) error {
	query := `INSERT INTO refresh_tokens (guid, hashed_token, created_at) VALUES ($1, $2, $3)`
	_, err := s.DB.Exec(query, guid, hashToken, time.Now())
	if err != nil {
		return err
	}
	return nil
}

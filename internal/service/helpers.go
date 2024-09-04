package service

import "auth/internal/config"

func (s *service) getSecretKey(cfg config.JWT) []byte {
	return []byte(cfg.SecretKey)
}

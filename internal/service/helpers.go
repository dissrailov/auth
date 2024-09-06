package service

import (
	"auth/internal/config"
	"fmt"
)

func (s *service) getSecretKey(cfg config.JWT) []byte {
	return []byte(cfg.SecretKey)
}

func (s *service) sendEmail(email, oldIP, newIP string) error {
	fmt.Printf("Sending email to %s\n Dear user your ip has changed from %s to %s\n", email, oldIP, newIP)
	return nil
}

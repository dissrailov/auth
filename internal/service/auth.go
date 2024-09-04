package service

import (
	"auth/internal/config"
	"auth/internal/models"
	"crypto/rand"
	"encoding/base64"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) ForgeAuthPair(guid string, ip string, cfg config.Config) (*models.TokenPair, error) {
	accessToken := s.generateAccessToken(ip)
	accessTokenSecret, err := accessToken.SignedString(s.getSecretKey(cfg.JWT))
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}
	err = s.saveRefreshToken(guid, refreshToken, ip)
	if err != nil {
		return nil, err
	}
	pair := &models.TokenPair{
		AccessObject: accessToken,
		AccessToken:  accessTokenSecret,
		RefreshToken: refreshToken,
	}
	return pair, nil
}

func (s *service) generateRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

func (s *service) generateAccessToken(ip string) *jwt.Token {
	claims := jwt.MapClaims{
		"ip": ip,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token
}

func (s *service) saveRefreshToken(guid string, refreshToken string, ip string) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = s.repo.InsertRefreshToken(guid, string(hashedToken), ip)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) VerifyRefreshToken(refreshToken, guid, currentIP string) (string, bool, error) {
	hashedToken, ipDB, err := s.repo.GetRefreshToken(guid)
	if err != nil {
		return "", false, err
	}
	if ipDB != currentIP {

	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(refreshToken))
	if err != nil {
		return "", false, err
	}

	return guid, true, nil
}

func (s *service) DeleteRefreshToken(guid string) error {
	err := s.repo.DeleteRefreshToken(guid)
	if err != nil {
		return err
	}
	return nil
}

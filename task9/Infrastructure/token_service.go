package infrastructure

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type tokenService struct {}

func NewTokenService() *tokenService {
	return &tokenService{}
}

func (t *tokenService) GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (t *tokenService) HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
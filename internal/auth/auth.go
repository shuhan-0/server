package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"snake-game/internal/cache"
	"snake-game/internal/database"
	"time"
)

type AuthService struct {
	db    *database.MySQLDB
	cache *cache.RedisClient
}

func NewAuthService(db *database.MySQLDB, cache *cache.RedisClient) *AuthService {
	return &AuthService{
		db:    db,
		cache: cache,
	}
}

func (as *AuthService) Authenticate(username, password string) (string, error) {
	// In a real application, you would check the username and password against the database
	// For this example, we'll just generate a token

	token, err := generateToken()
	if err != nil {
		return "", err
	}

	// Store the token in Redis with an expiration time
	err = as.cache.Client.Set(token, username, time.Hour).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) ValidateToken(token string) (string, error) {
	username, err := as.cache.Client.Get(token).Result()
	if err != nil {
		return "", errors.New("invalid token")
	}
	return username, nil
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

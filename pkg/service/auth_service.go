package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
)

type AuthService interface {
	Register(user model.User) error
	Login(username, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(user model.User) error {
	return s.repo.Register(user)
}

func (s *authService) Login(username, password string) (string, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}

	// Verify password here

	accessToken, err := generateToken(strconv.FormatUint(uint64(user.ID), 10), "access", 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(strconv.FormatUint(uint64(user.ID), 10), "refresh", 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
	claims, err := validateToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	if claims["type"] != "refresh" {
		return "", "", errors.New("invalid token type")
	}

	userID := claims["user_id"].(string)

	newAccessToken, err := generateToken(userID, "access", 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := generateToken(userID, "refresh", 7*24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func generateToken(userID string, tokenType string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"exp":     time.Now().Add(expirationTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key")) // Replace with a secure secret key
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // Replace with the same secret key used for signing
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

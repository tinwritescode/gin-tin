package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tinwritescode/gin-tin/pkg/config"
	"github.com/tinwritescode/gin-tin/pkg/model"
	"github.com/tinwritescode/gin-tin/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user model.User) error
	Login(username, password string) (string, string, error)
	RefreshToken(refreshToken string) (string, string, error)
}

type authService struct {
	repo   repository.AuthRepository
	config *config.Config
}

func NewAuthService(repo repository.AuthRepository, cfg *config.Config) AuthService {
	return &authService{repo: repo, config: cfg}
}

func (s *authService) Register(user model.User) error {
	if user.Role == "" {
		user.Role = model.RoleUser // Default role
	}
	return s.repo.Register(user)
}

func (s *authService) Login(username, password string) (string, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}

	// Verify password here
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", err
	}

	accessToken, err := s.generateToken(strconv.FormatUint(uint64(user.ID), 10), "access", 15*time.Minute, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.generateToken(strconv.FormatUint(uint64(user.ID), 10), "refresh", 7*24*time.Hour, user.Role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	if claims["type"] != "refresh" {
		return "", "", errors.New("invalid token type")
	}

	userID := claims["user_id"].(string)

	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return "", "", err
	}

	newAccessToken, err := s.generateToken(userID, "access", 15*time.Minute, user.Role)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := s.generateToken(userID, "refresh", 7*24*time.Hour, user.Role)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *authService) generateToken(userID string, tokenType string, expirationTime time.Duration, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"exp":     time.Now().Add(expirationTime).Unix(),
		"role":    role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWTSecretKey))
}

func (s *authService) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

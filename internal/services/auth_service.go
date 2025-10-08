package services

import (
	"errors"
	"meeting_scheduler/internal/repository"
	"meeting_scheduler/internal/utils"
	"time"
)

type AuthService struct {
	users     repository.UserRepo
	jwtSecret string
}

func NewAuthService(users repository.UserRepo, jwtSecret string) *AuthService {
	return &AuthService{users: users, jwtSecret: jwtSecret}
}

// Authenticate checks credentials and returns a JWT token.
func (s *AuthService) Authenticate(email, password string) (string, error) {
	u, err := s.users.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if err := utils.CheckPassword(u.PasswordHash, password); err != nil {
		return "", errors.New("invalid credentials")
	}
	token, err := utils.GenerateToken(s.jwtSecret, u.ID, u.IsAdmin, 24*time.Hour)
	return token, err
}

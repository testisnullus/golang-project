package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/testisnullus/golang-project/pkg/config"
	"github.com/testisnullus/golang-project/pkg/jwt"
	"github.com/testisnullus/golang-project/pkg/models"
	"github.com/testisnullus/golang-project/pkg/users/repository"
	"math/rand"
)

type AuthService struct {
	repo repository.Authorization
	cfg  *config.YamlFile
}

func NewAuthService(repo repository.Authorization, cfg *config.YamlFile) *AuthService {
	return &AuthService{repo: repo, cfg: cfg}
}

func (s *AuthService) CreateUser(ctx context.Context, user *models.User) error {
	salt := saltGeneration(20)

	err := s.repo.CreateUser(ctx, user, salt, hashPassword(salt, user.Password))
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(salt, password string) string {
	password = salt + password
	hashBits := sha256.Sum256([]byte(password))
	hash := base64.StdEncoding.EncodeToString(hashBits[:])

	return hash
}

func saltGeneration(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (s *AuthService) Login(ctx context.Context, user *models.LoginUser) (string, error) {
	userID, err := s.repo.GetUserByEmail(ctx, user)
	if err != nil {
		return "", err
	}

	salt, password, err := s.repo.GetPasswordByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	hashFrontPassword := hashPassword(salt, user.Password)

	if hashFrontPassword == password {
		token, err := jwt.NewJWT(userID, []byte(s.cfg.JWT.HmacSecret), s.cfg.JWT.Lifetime)
		if err != nil {
			return "", err
		}
		return token, nil
	}

	return "", fmt.Errorf("Invalid data")
}

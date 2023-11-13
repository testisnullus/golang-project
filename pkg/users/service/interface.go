package service

import (
	"context"
	"github.com/testisnullus/golang-project/pkg/config"
	"github.com/testisnullus/golang-project/pkg/models"
	"github.com/testisnullus/golang-project/pkg/users/repository"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *models.User) error
	Login(ctx context.Context, user *models.LoginUser) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, cfg *config.YamlFile) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, cfg),
	}
}

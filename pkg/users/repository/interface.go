package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/testisnullus/golang-project/pkg/models"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *models.User, salt, password string) error
	GetPasswordByUserID(ctx context.Context, userID uint64) (string, string, error)
	GetUserByEmail(ctx context.Context, user *models.LoginUser) (uint64, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}

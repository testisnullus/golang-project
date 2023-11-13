package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/testisnullus/golang-project/pkg/models"
)

func NewAuthRepository(db *sqlx.DB) *DB {
	return &DB{db}
}

func (r *DB) CreateUser(ctx context.Context, user *models.User, salt, password string) error {
	tx, err := r.Begintx(ctx)
	if err != nil {
		tx.Rollback()
		return err
	}

	var id int
	err = tx.QueryRowContext(ctx, "INSERT INTO users (first_name, last_name, email) VALUES ($1,$2,$3) RETURNING id", user.FirstName, user.LastName, user.Email).Scan(&id)
	if err != nil {
		return err
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO hash_password (user_id, salt, password) VALUES ($1,$2,$3)", id, salt, password)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (r *DB) GetUserByEmail(ctx context.Context, user *models.LoginUser) (uint64, error) {
	tx, err := r.Begintx(ctx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var userID uint64
	err = r.QueryRowContext(ctx, "SELECT id from users where email=$1", user.Email).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *DB) GetPasswordByUserID(ctx context.Context, userID uint64) (string, string, error) {
	tx, err := r.Begintx(ctx)
	if err != nil {
		tx.Rollback()
		return "", "", err
	}

	var salt string
	var hashPassword string
	err = r.QueryRowContext(ctx, "SELECT salt, password from hash_password where user_id=$1", userID).Scan(&salt, &hashPassword)
	if err != nil {
		return "", "", err
	}

	return salt, hashPassword, err
}

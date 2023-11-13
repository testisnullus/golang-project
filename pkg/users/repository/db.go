package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func (db *DB) Begintx(ctx context.Context) (*sqlx.Tx, error) {

	tx, err := db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

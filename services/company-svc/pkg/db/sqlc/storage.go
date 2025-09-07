package db

import (
	"context"
	"database/sql"
	"fmt"
)

type SQLStorage struct {
	*Queries
	db *sql.DB
}

func (storage *SQLStorage) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := storage.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func NewStorage(db *sql.DB) *SQLStorage {
	return &SQLStorage{
		db:      db,
		Queries: New(db),
	}
}

package db

import (
	"context"
	"database/sql"
	"errors"
)

type Store struct {
	Querier
	db *sql.DB
}

// NewStore creates a new store for the given database
func NewStore(db *sql.DB) *Store {
	return &Store{
		Querier: New(db),
		db:      db,
	}
}

// ExecTx executes a function within a database transaction
func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rBErr := tx.Rollback(); rBErr != nil {
			return errors.New("rollback error")
		}
		return err
	}

	return tx.Commit()
}

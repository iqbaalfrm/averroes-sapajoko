package postgres

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka koneksi database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("gagal ping database: %w", err)
	}
	return db, nil
}

func (r *Repository) DB() *sql.DB {
	return r.db
}

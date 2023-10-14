package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	_ "github.com/jackc/pgx/v5/stdlib"
	//_ "github.com/mattn/go-sqlite3"
)

func NewDB(cnf *config.Config) (DB, error) {
	db, err := sql.Open("pgx", cnf.DatabaseDsn)
	if err != nil {
		return DB{}, fmt.Errorf("can't connect to database: %w", err)
	}
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS urli (
			id SERIAL PRIMARY KEY,
			short_url TEXT NOT NULL,
			original_url TEXT NOT NULL
		)
	`
	_, err = db.ExecContext(context.Background(), createTableQuery)
	if err != nil {
		return DB{}, fmt.Errorf("can't create table to database: %w", err)
	}
	return DB{db: db}, nil

}

type DB struct {
	db *sql.DB
}

func (d *DB) Get(key string) (string, error) {
	var res string
	err := d.db.QueryRowContext(context.Background(), "SELECT * FROM urls WHERE short_url = ?", key).Scan(&res)
	if err != nil {
		return "", fmt.Errorf("can't get url: %w", err)
	}

	return res, nil
}

func (d *DB) Set(key string, value string) error {
	query := "INSERT INTO urls (original_url, short_url) VALUES (?, ?) " +
		"ON DUPLICATE KEY UPDATE " +
		"original_url = VALUES(original_url)"
	_, err := d.db.ExecContext(context.Background(), query, key, value)
	if err != nil {
		return consts.ErrGetURL
	}
	return nil
}

func (d *DB) Count() int {
	row := d.db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM urls")
	var res int
	err := row.Scan(&res)
	if err != nil {
		return 0
	}
	return res

}

func (d *DB) Ping() error {
	if err := d.db.Ping(); err != nil {
		return fmt.Errorf("ping error: %w", err)
	}
	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

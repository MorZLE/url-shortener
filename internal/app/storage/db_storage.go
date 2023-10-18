package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(cnf *config.Config) (DB, error) {

	db, err := sql.Open("pgx", cnf.DatabaseDsn)
	if err != nil {
		return DB{}, fmt.Errorf("can't connect to database: %w", err)
	}
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS urls (
		    id SERIAL PRIMARY KEY,
			short_url TEXT UNIQUE,
			original_url TEXT UNIQUE
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
	err := d.db.QueryRowContext(context.Background(), "SELECT original_url FROM urls WHERE short_url = $1", key).Scan(&res)
	if err != nil {
		return "", fmt.Errorf("can't get url: %w", err)
	}

	return res, nil
}

func (d *DB) Set(key string, value string) (string, error) {

	query := `INSERT INTO urls (short_url, original_url) ON CONFLICT (long_url_url) VALUES ($1, $2)`
	_, err := d.db.ExecContext(context.Background(), query, key, value)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return d.GetDuplicate(value)
			}
		}
		if errors.Is(err, consts.ErrKeyBusy) {
			return "", consts.ErrKeyBusy
		}
	}

	return "", nil
}

func (d *DB) GetDuplicate(longURL string) (string, error) {
	var value string
	query := `SELECT short_url FROM urls WHERE original_url = $1`
	err := d.db.QueryRowContext(context.Background(), query, longURL).Scan(&value)
	if err != nil {
		return "", fmt.Errorf("can't get dublicate url: %w", err)
	}
	return value, consts.ErrDuplicateURL
}

func (d *DB) SetBatch(m map[string]string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES ($1, $2)`
	tr, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("can't start transaction: %w", err)
	}
	for key, value := range m {
		_, err = tr.ExecContext(context.Background(), query, key, value)
		if err != nil {
			tr.Rollback()
			return fmt.Errorf("can't set url: %w", err)
		}
	}
	err = tr.Commit()
	if err != nil {
		return fmt.Errorf("can't commit transaction: %w", err)
	}
	return nil
}
func (d *DB) Count() int {
	var res int
	err := d.db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM urls").Scan(&res)
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

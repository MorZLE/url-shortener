package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/MorZLE/url-shortener/internal/config"
	"github.com/MorZLE/url-shortener/internal/consts"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
)

const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS urls (
			short_url TEXT UNIQUE,
			original_url TEXT UNIQUE,
            user_id TEXT 
		)`
	insertQuery       = `INSERT INTO urls (short_url, original_url, user_id) VALUES ($1, $2, $3)`
	selectOriginalURL = `SELECT original_url FROM urls WHERE short_url = $1 `
	selectShortURL    = `SELECT short_url FROM urls WHERE original_url = $1 `
	selectCount       = `SELECT COUNT(*) FROM urls`
	selectAllUsersURL = `SELECT short_url, original_url FROM urls WHERE user_id = $1`
)

func NewDB(cnf *config.Config) (DB, error) {
	db, err := sql.Open("postgres", cnf.DatabaseDsn)
	if err != nil {
		return DB{}, fmt.Errorf("can't connect to database: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return DB{}, fmt.Errorf("failed to create migrate driver, %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"url", driver)
	if err != nil {
		return DB{}, fmt.Errorf("failed to migrate: %w", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return DB{}, fmt.Errorf("failed to do migrate %w", err)
	}

	return DB{db: db}, nil

}

type DB struct {
	db *sql.DB
}

func (d *DB) Get(key string) (string, error) {
	var res string
	err := d.db.QueryRowContext(context.Background(), selectOriginalURL, key).Scan(&res)
	if err != nil {
		return "", fmt.Errorf("can't get url: %w", err)
	}

	return res, nil
}

func (d *DB) Set(id, key, value string) error {
	err := d.db.QueryRowContext(context.Background(), insertQuery, key, value, id).Err()
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return consts.ErrDuplicateURL
			}
		}
		return fmt.Errorf("can't set url: %w", err)
	}
	return nil
}

func (d *DB) GetDuplicate(longURL string) (string, error) {
	var value string
	err := d.db.QueryRowContext(context.Background(), selectShortURL, longURL).Scan(&value)
	if err != nil {
		return "", fmt.Errorf("can't get dublicate url: %w", err)
	}
	return value, nil
}

func (d *DB) SetBatch(id string, m map[string]string) error {
	tr, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("can't start transaction: %w", err)
	}
	for key, value := range m {
		_, err = tr.ExecContext(context.Background(), insertQuery, key, value, id)
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

func (d *DB) GetAllURL(id string) (map[string]string, error) {
	m := make(map[string]string)
	rows, err := d.db.QueryContext(context.Background(), selectAllUsersURL, id)
	if err != nil {
		return m, fmt.Errorf("can't get all urls: %w", err)
	}
	for rows.Next() {
		err := rows.Err()
		if err != nil {
			return m, fmt.Errorf("can't get all urls: %w", err)

		}
		var key, value string
		err = rows.Scan(&key, &value)
		if err != nil {
			return m, fmt.Errorf("can't change map urls: %w", err)
		}
		m[key] = value
	}
	return m, nil
}

func (d *DB) Count() (int, error) {
	var res int
	err := d.db.QueryRowContext(context.Background(), selectCount).Scan(&res)
	if err != nil {
		return 0, fmt.Errorf("can't count urls: %w", err)
	}
	return res, nil

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

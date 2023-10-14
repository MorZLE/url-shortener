package storage

import (
	"github.com/MorZLE/url-shortener/internal/config"
	//"github.com/MorZLE/url-shortener/internal/consts"
	//"github.com/MorZLE/url-shortener/internal/constjson"
	//"log"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB(cnf *config.Config) (DB, error) {

	db, err := sql.Open("pgx", cnf.ServerAddr)
	if err != nil {
		return DB{}, err
	}
	return DB{db: db}, nil

}

type DB struct {
	db *sql.DB
}

func (d *DB) Get(key string) (string, error) {
	return "", nil
}

func (d *DB) Set(key string, value string) error {
	return nil
}

func (d *DB) Count() int {
	return 0

}

func (d *DB) Ping() error {
	if err := d.db.Ping(); err != nil {
		return err
	}
	return nil
}
func (d *DB) Close() error {
	return d.db.Close()
}

package postgres

import (
	"database/sql"
	"errors"

	"github.com/bekryasheva/url-shortener/pkg"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase(db *sql.DB) *PostgresDatabase {
	return &PostgresDatabase{db: db}
}

func (p *PostgresDatabase) Save(originalURL string) (int64, error) {
	var id int64
	err := p.db.QueryRow(`INSERT INTO urls (url) VALUES ($1) ON CONFLICT (url) DO UPDATE SET url=excluded.url RETURNING id`, originalURL).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PostgresDatabase) Get(id int64) (string, error) {
	var url string

	err := p.db.QueryRow("SELECT url FROM urls where id = $1", id).Scan(&url)
	if errors.Is(err, sql.ErrNoRows) {
		return "", pkg.ErrNotFound
	}

	if err != nil {
		return "", err
	}

	return url, nil
}

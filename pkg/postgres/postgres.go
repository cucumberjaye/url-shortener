package postgres

import (
	"database/sql"
	"github.com/cucumberjaye/url-shortener/configs"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New() (*sql.DB, error) {
	db, err := sql.Open("pgx", configs.DataBaseDSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}

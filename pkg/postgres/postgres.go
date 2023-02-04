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

	err = createTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
    	user_id integer not null,
    	short_url varchar(255) not null,
    	original_url varchar(255) not null unique,
    	uses integer not null)`

	_, err := db.Exec(query)

	return err
}

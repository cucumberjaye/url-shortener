package postrgres

import "database/sql"

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (r *DB) CheckDBConn() error {
	return r.db.Ping()
}

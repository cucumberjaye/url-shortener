package postrgresdb

import (
	"database/sql"
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (k *SQLStore) CheckKeeper() error {
	return k.db.Ping()
}

func (k *SQLStore) GetAllData() (repository.DB, error) {
	users := repository.DB{
		Store: map[string]map[string]string{},
		Exist: map[string]map[string]int{},
	}

	query := "SELECT * FROM urls"
	row, err := k.db.Query(query)
	if err != nil {
		return users, err
	}
	defer row.Close()

	for row.Next() {
		var id, short, full string
		var count int
		var deleted bool
		if err = row.Scan(&id, &short, &full, &count, &deleted); err != nil {
			return users, err
		}
		if deleted {
			count = -1
		}
		if _, ok := users.Store[id]; !ok {
			users.Store[id] = map[string]string{short: full}
			users.Exist[id] = map[string]int{full: count}
		} else {
			users.Store[id][short] = full
			users.Exist[id][full] = count
		}
	}

	if err = row.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func (k *SQLStore) Set(users repository.DB) error {
	for key, val := range users.Store {
		for short, full := range val {
			if short != "" {
				query := "INSERT INTO urls (user_id, short_url, original_url, uses, deleted) VALUES ($1, $2, $3, $4, $5)"
				_, err := k.db.Exec(query, key, short, full, 0, false)
				if err != nil {
					return err
				}
				return nil
			}
		}
	}

	for _, val := range users.Exist {
		for full, count := range val {
			if count == -1 {
				query := "UPDATE urls SET deleted=TRUE WHERE original_url=$1"
				_, err := k.db.Exec(query, full)
				if err != nil {
					return err
				}
			} else {
				query := "UPDATE urls SET uses=$1 WHERE original_url=$2"
				_, err := k.db.Exec(query, count, full)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}

	return nil
}

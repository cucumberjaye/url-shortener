package postrgresdb

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cucumberjaye/url-shortener/models"
)

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db: db,
	}
}

func (r *SQLStore) SetURL(fullURL, shortURL string, id int) error {
	selectQuery := "SELECT COUNT(*) FROM urls WHERE original_url=$1"
	row, err := r.db.Query(selectQuery, fullURL)
	if err != nil {
		return err
	}
	defer row.Close()
	var count int
	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		query := "INSERT INTO urls (user_id, short_url, original_url, uses) values ($1, $2, $3, $4)"
		_, err = r.db.Exec(query, id, shortURL, fullURL, 0)
		if err != nil {
			return err
		}
	} else {
		return errors.New("url already exists")
	}

	return row.Err()
}

func (r *SQLStore) GetURL(shortURL string) (string, error) {
	query := "SELECT original_url FROM urls WHERE short_url=$1"
	row, err := r.db.Query(query, shortURL)
	if err != nil {
		return "", err
	}

	defer row.Close()

	var fullURL string
	for row.Next() {
		if err = row.Scan(&fullURL); err != nil {
			return "", err
		}
	}

	if err = row.Err(); err != nil {
		return "", err
	}

	updateQuery := "UPDATE urls SET uses=(SELECT uses FROM urls WHERE short_url=$1)+1 WHERE short_url=$1"
	_, err = r.db.Exec(updateQuery, shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, nil
}

func (r *SQLStore) GetURLCount() (int64, error) {
	query := "SELECT COUNT(*) FROM urls"
	row, err := r.db.Query(query)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	var count int64
	for row.Next() {
		if err = row.Scan(&count); err == sql.ErrNoRows {
			fmt.Println("check")
			return 0, nil
		} else if err != nil {
			return 0, err
		}
	}

	if err = row.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *SQLStore) GetAllUserURL(id int) ([]models.URLs, error) {
	query := "SELECT short_url, original_url FROM urls WHERE user_id=$1"
	row, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	fmt.Println(id)

	var out = []models.URLs{}
	for row.Next() {
		var v = models.URLs{}
		err = row.Scan(&v.ShortURL, &v.OriginalURL)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}

	err = row.Err()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *SQLStore) GetRequestCount(shortURL string) (int, error) {
	query := "SELECT uses FROM urls WHERE short_url=$1"

	row, err := r.db.Query(query, shortURL)
	if err != nil {
		return 0, err
	}

	defer row.Close()

	var count int

	for row.Next() {
		if err = row.Scan(&count); err != nil {
			return 0, err
		}
	}

	if err = row.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *SQLStore) CheckDBConn() error {
	return r.db.Ping()
}

func (r *SQLStore) BatchSetURL(data []models.BatchInputJSON, shortURL []string, id int) ([]models.BatchInputJSON, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	selStmt, err := tx.Prepare("SELECT COUNT(*) FROM urls WHERE original_url=$1")
	if err != nil {
		return nil, err
	}
	defer selStmt.Close()

	insStmt, err := tx.Prepare("INSERT INTO urls (user_id, short_url, original_url, uses) values ($1, $2, $3, $4)")
	if err != nil {
		return nil, err
	}
	defer insStmt.Close()

	for i := 0; i < len(data); i++ {
		row, err := selStmt.Query(data[i].OriginalURL)
		if err != nil {
			return nil, err
		}
		var count int
		row.Next()
		err = row.Scan(&count)
		if err != nil {
			return nil, err
		}

		if err = row.Err(); err != nil {
			return nil, err
		}
		row.Close()

		if count == 0 {
			_, err = insStmt.Exec(id, shortURL[i], data[i].OriginalURL, 0)
			if err != nil {
				if err = tx.Rollback(); err != nil {
					return nil, err
				}
			}
			data[i].OriginalURL = shortURL[i]
		} else {
			return nil, errors.New("url already exists")
		}
	}

	return data, tx.Commit()

}

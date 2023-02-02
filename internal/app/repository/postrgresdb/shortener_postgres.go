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
	selectQuery := "SELECT short_url FROM urls WHERE short_url=$1"
	row, err := r.db.Query(selectQuery, shortURL)
	if err != nil {
		return err
	}
	defer row.Close()
	var short string
	row.Next()
	err = row.Scan(&short)
	if err == sql.ErrNoRows {
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

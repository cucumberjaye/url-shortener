package app

import "errors"

type Database struct {
	Store map[string]string
}

func NewDB() *Database {
	store := make(map[string]string)
	return &Database{Store: store}
}

func (d *Database) SetURL(fullURL, shortURL string) error {
	if _, ok := d.Store[shortURL]; !ok {
		d.Store[shortURL] = fullURL
		return nil
	} else {
		return errors.New("url already exist")
	}
}

func (d *Database) GetURL(shortURL string) (string, error) {
	if url, ok := d.Store[shortURL]; !ok {
		return "", errors.New("url is not exist")
	} else {
		return url, nil
	}
}

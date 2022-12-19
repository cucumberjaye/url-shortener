package app

import "errors"

type Database struct {
	Store map[string]string
}

func NewDB() *Database {
	store := make(map[string]string)
	return &Database{Store: store}
}

func (d *Database) SetUrl(fullUrl, shortUrl string) error {
	if _, ok := d.Store[shortUrl]; !ok {
		d.Store[shortUrl] = fullUrl
		return nil
	} else {
		return errors.New("url already exist")
	}
}

func (d *Database) GetUrl(shortUrl string) (string, error) {
	if url, ok := d.Store[shortUrl]; !ok {
		return "", errors.New("url is not exist")
	} else {
		return url, nil
	}
}

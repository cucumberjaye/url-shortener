package mocks

import (
	"errors"
)

type RepositoryMock struct {
}

func (r *RepositoryMock) SetURL(fullURL, shortURL string) error {
	if fullURL == "test.com" {
		return nil
	}
	return errors.New("test")
}

func (r *RepositoryMock) GetURL(shortURL string) (string, error) {
	if shortURL == "0" {
		return "test.com", nil
	}
	return "", errors.New("test")
}

func (r *RepositoryMock) GetURLCount() int64 {
	return 0
}

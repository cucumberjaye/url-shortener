package mocks

import (
	"errors"
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
)

type RepositoryMock struct {
	repository.Shortener
}

func (r *RepositoryMock) SetURL(fullURL, shortURL string) error {
	if fullURL == "test.com" {
		return nil
	}
	return errors.New("test")
}

func (r *RepositoryMock) GetURL(shortURL string) (string, error) {
	if shortURL == "test" {
		return "test.com", nil
	}
	return "", errors.New("test")
}

package mocks

import (
	"errors"

	"github.com/cucumberjaye/url-shortener/models"
)

// мок структура для репозитория
type RepositoryMock struct {
}

// мок для SetURL
func (r *RepositoryMock) SetURL(fullURL, shortURL string, id int) error {
	if fullURL == "test.com" {
		return nil
	}
	return errors.New("test")
}

// мок для GetURL
func (r *RepositoryMock) GetURL(shortURL string) (string, error) {
	if shortURL == "0" {
		return "test.com", nil
	}
	return "", errors.New("test")
}

// мок для GetURLCount
func (r *RepositoryMock) GetURLCount() int64 {
	return 0
}

// мок для GetAllUserURL
func (r *RepositoryMock) GetAllUserURL(id int) []models.URLs {
	return []models.URLs{}
}

package mocks

import (
	"errors"

	"github.com/cucumberjaye/url-shortener/models"
)

// мок для структуры ShortenerService
type ServiceMock struct {
}

// мок для ShortingURL
func (m *ServiceMock) ShortingURL(fullURL, baseURL string, id string) (string, error) {
	if fullURL == "test.com" {
		return "0", nil
	}
	return "", errors.New("test")
}

// мок для GetFullURL
func (m *ServiceMock) GetFullURL(shortURL string) (string, error) {
	if shortURL[len(shortURL)-1] == '0' {
		return "test.com", nil
	}
	return "", errors.New("test")
}

// мок для GetAllUserURL
func (m *ServiceMock) GetAllUserURL(id string) ([]models.URLs, error) {
	return []models.URLs{}, nil
}

// мок для CheckDBConn
func (m *ServiceMock) CheckDBConn() error {
	return nil
}

// мок для BatchSetURL
func (m *ServiceMock) BatchSetURL(data []models.BatchInputJSON, baseURL string, id string) ([]models.BatchInputJSON, error) {
	return []models.BatchInputJSON{}, nil
}

// мок для BatchDeleteURL
func (m *ServiceMock) BatchDeleteURL(data []string, id string) {
}

// мок для GetStats
func (m *ServiceMock) GetStats() (models.Stats, error) {
	return models.Stats{}, nil
}

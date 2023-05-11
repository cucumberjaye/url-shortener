package service

import (
	"github.com/cucumberjaye/url-shortener/models"
)

// Интерфейс хранилища для логирования
type LogsInfoRepository interface {
	GetRequestCount(shortURL string) (int, error)
}

// интерфейс хранилища для сокращателя ссылок
type URLRepository interface {
	SetURL(fullURL, shortURL string, id string) (string, error)
	GetURL(shortURL string) (string, error)
	GetURLCount() (int64, error)
	GetAllUserURL(id string) ([]models.URLs, error)
	BatchSetURL(data []models.BatchInputJSON, shortURL []string, id string) ([]models.BatchInputJSON, error)
	CheckStorage() error
	GetStats() (models.Stats, error)
}

// общий интерфейс
type URLLogs interface {
	URLRepository
	LogsInfoRepository
}

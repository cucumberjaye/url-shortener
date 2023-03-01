package service

import (
	"github.com/cucumberjaye/url-shortener/models"
)

type LogsInfoRepository interface {
	GetRequestCount(shortURL string) (int, error)
}

type URLRepository interface {
	SetURL(fullURL, shortURL string, id string) (string, error)
	GetURL(shortURL string) (string, error)
	GetURLCount() (int64, error)
	GetAllUserURL(id string) ([]models.URLs, error)
	BatchSetURL(data []models.BatchInputJSON, shortURL []string, id string) ([]models.BatchInputJSON, error)
	CheckStorage() error
}

type URLLogs interface {
	URLRepository
	LogsInfoRepository
}

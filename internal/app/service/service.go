package service

import (
	"github.com/cucumberjaye/url-shortener/models"
)

type LogsInfoRepository interface {
	GetRequestCount(shortURL string) (int, error)
}

type URLRepository interface {
	SetURL(fullURL, shortURL string, id int) error
	GetURL(shortURL string) (string, error)
	GetURLCount() int64
	GetAllUserURL(id int) []models.URLs
}

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
	GetURLCount() (int64, error)
	GetAllUserURL(id int) ([]models.URLs, error)
}

type SQLRepository interface {
	CheckDBConn() error
}

type URLLogs interface {
	URLRepository
	LogsInfoRepository
}

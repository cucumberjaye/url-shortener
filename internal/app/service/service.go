package service

import (
	"github.com/cucumberjaye/url-shortener/models"
)

type LogsInfoRepository interface {
	GetRequestCount(shortURL string, id int) (int, error)
}

type URLRepository interface {
	SetURL(fullURL, shortURL string, id int) error
	GetURL(shortURL string, id int) (string, error)
	GetURLCount() map[int]int
	GetAllUserURL(id int) []models.URLs
}

package service

import (
	"github.com/cucumberjaye/url-shortener/internal/app/repository"
)

type Shortener interface {
	ShortingURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

type Service struct {
	Shortener
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Shortener: NewShortenerService(repos),
	}
}

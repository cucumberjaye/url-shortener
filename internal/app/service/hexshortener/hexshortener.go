package hexshortener

import (
	"fmt"
	"github.com/cucumberjaye/url-shortener/internal/app/service"
	"github.com/cucumberjaye/url-shortener/models"
	"sync"
)

type ShortenerService struct {
	repos   service.URLRepository
	counter map[int]int
	mx      sync.Mutex
}

func NewShortenerService(repos service.URLRepository) *ShortenerService {
	return &ShortenerService{
		repos:   repos,
		counter: repos.GetURLCount(),
	}
}

func (s *ShortenerService) ShortingURL(fullURL, baseURL string, id int) (string, error) {
	shortURL := baseURL + fmt.Sprintf("%x", s.counter[id])
	if err := s.repos.SetURL(fullURL, shortURL, id); err != nil {
		return "", err
	}
	s.mx.Lock()
	s.counter[id]++
	s.mx.Unlock()

	return shortURL, nil
}

func (s *ShortenerService) GetFullURL(shortURL string, id int) (string, error) {
	fullURL, err := s.repos.GetURL(shortURL, id)
	if err != nil {
		return "", err
	}

	return fullURL, err
}

func (s *ShortenerService) GetAllUserURL(id int) []models.URLs {
	return s.repos.GetAllUserURL(id)
}

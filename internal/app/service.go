package app

import "math/rand"

type Service struct {
	db *Database
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewService(db *Database) *Service {
	return &Service{db: db}
}

func (s *Service) ShortingURL(fullURL string) (string, error) {
	shortURL := shorting()
	if err := s.db.SetURL(fullURL, shortURL); err != nil {
		return "", err
	}

	return "http://localhost:8080/" + shortURL, nil
}

func (s *Service) GetFullURL(shortURL string) (string, error) {
	fullURL, err := s.db.GetURL(shortURL)
	if err != nil {
		return "", err
	}

	return fullURL, err
}

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

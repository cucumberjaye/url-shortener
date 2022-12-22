package repository

type Shortener interface {
	SetURL(fullURL, shortURL string) error
	GetURL(shortURL string) (string, error)
}

type Repository struct {
	Shortener
}

func NewRepository() *Repository {
	return &Repository{Shortener: NewShortenerDB()}
}

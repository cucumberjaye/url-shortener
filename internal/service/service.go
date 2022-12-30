package service

type LogsInfoRepository interface {
	GetRequestCount(shortURL string) (int, error)
}

type URLRepository interface {
	SetURL(fullURL, shortURL string) error
	GetURL(shortURL string) (string, error)
}

package shortenerlogsinfo

import (
	"github.com/cucumberjaye/url-shortener/internal/app/service"
)

// струкура для логирования запросов
type URLLogsInfo struct {
	repos service.LogsInfoRepository
}

// создаем URLLogsInfo
func NewURLLogsInfo(repos service.LogsInfoRepository) *URLLogsInfo {
	return &URLLogsInfo{repos: repos}
}

// получаем количество получений по сокращенной ссылке
func (s *URLLogsInfo) GetRequestCount(shortURL string) (int, error) {
	count, err := s.repos.GetRequestCount(shortURL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

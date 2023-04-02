package shortenerlogsinfo

import (
	"github.com/cucumberjaye/url-shortener/internal/app/service"
)

type URLLogsInfo struct {
	repos service.LogsInfoRepository
}

func NewURLLogsInfo(repos service.LogsInfoRepository) *URLLogsInfo {
	return &URLLogsInfo{repos: repos}
}

func (s *URLLogsInfo) GetRequestCount(shortURL string) (int, error) {
	count, err := s.repos.GetRequestCount(shortURL)
	if err != nil {
		return 0, err
	}
	return count, nil
}

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

func (s *URLLogsInfo) GetRequestCount(shortURL string, id int) (int, error) {
	count, err := s.repos.GetRequestCount(shortURL, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}

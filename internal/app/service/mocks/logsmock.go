package mocks

// мок для структуры логов
type LogsMock struct {
}

// мок для GetRequestCount
func (s *LogsMock) GetRequestCount(shortURL string) (int, error) {
	return 0, nil
}

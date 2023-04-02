package mocks

type LogsMock struct {
}

func (s *LogsMock) GetRequestCount(shortURL string) (int, error) {
	return 0, nil
}

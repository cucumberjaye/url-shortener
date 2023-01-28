package mocks

type LogsMock struct {
}

func (s *LogsMock) GetRequestCount(shortURL string, id int) (int, error) {
	return 0, nil
}

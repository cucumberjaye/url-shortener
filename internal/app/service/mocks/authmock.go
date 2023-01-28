package mocks

type AuthMock struct {
}

func (s *AuthMock) GenerateNewToken() (string, error) {
	return "", nil
}

func (s *AuthMock) CheckToken(token string) (int, error) {
	return 0, nil
}

func (s *AuthMock) SetCurrentID(id int) {
}

func (s *AuthMock) GetCurrentID() int {
	return 0
}

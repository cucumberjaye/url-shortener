package models

// хранит сокращенный и полный ссылки
type URLs struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// хранит полную ссылку и ее идентификатор для добавления пачкой
type BatchInputJSON struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

// ответ от сервера при добавлении пачкой
type BatchOutputJSON struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

// хранит данные для удаления ссылки
type DeleteData struct {
	ShortURL string
	ID       string
}

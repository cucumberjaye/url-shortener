package models

type URLs struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type BatchInputJSON struct {
	CorrelationId string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchOutputJSON struct {
	CorrelationId string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

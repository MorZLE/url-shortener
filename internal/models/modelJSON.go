package models

type URLLong struct {
	URL string `json:"url"`
}

type URLShort struct {
	Result string `json:"result"`
}

type URLFile struct {
	ShortURL    string `json:"result"`
	OriginalURL string `json:"original_url"`
}

type BatchSet struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type BatchGet struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

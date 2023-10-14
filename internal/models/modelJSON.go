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

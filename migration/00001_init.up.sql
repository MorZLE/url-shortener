CREATE TABLE IF NOT EXISTS urls (
    short_url TEXT UNIQUE,
    original_url TEXT UNIQUE
)
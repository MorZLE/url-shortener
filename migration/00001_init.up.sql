CREATE TABLE IF NOT EXISTS urls (
    short_url TEXT UNIQUE,
    original_url TEXT UNIQUE,
    user_id TEXT,
    delete_flag BOOLEAN default False
);
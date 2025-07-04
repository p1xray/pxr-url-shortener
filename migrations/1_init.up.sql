CREATE TABLE IF NOT EXISTS urls(
  id INTEGER PRIMARY KEY,
  long_url TEXT NOT NULL,
  short_code TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_urls_short_code ON urls(short_code);
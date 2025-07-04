package domain

import "time"

// URL is data for URL in storage.
type URL struct {
	ID        int64
	LongUrl   string
	ShortCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}

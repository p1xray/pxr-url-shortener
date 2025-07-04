package server

import "context"

// URLService represents a service for working with short URL.
type URLService interface {
	// Shorten generates a new unique short code for URL.
	Shorten(ctx context.Context, longURL string) (string, error)

	// LongURL returns a long URL by short code.
	LongURL(ctx context.Context, shortCode string) (string, error)
}

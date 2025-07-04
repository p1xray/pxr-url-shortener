package server

import "context"

// URLService represents a service for working with short URL.
type URLService interface {
	// LongURL returns a long URL by short code.
	LongURL(ctx context.Context, shortCode string) (string, error)
}

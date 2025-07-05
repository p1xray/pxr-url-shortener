package server

import (
	"context"
	"github.com/p1xray/pxr-url-shortener/internal/dto"
)

// URLService represents a service for working with short URL.
type URLService interface {
	// Shorten generates a new unique short code for URL.
	Shorten(ctx context.Context, longURL, host string) (dto.Shorten, error)

	// LongURL returns a long URL by short code.
	LongURL(ctx context.Context, shortCode string) (string, error)
}

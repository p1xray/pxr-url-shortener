package dto

// Shorten is an output DTO for the URL shortening service.
type Shorten struct {
	ShortCode string // Generated short code.
	ShortURL  string // Shortened URL.
}

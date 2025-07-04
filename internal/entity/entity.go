package entity

import (
	"errors"
	"fmt"
	shortcodegenerator "github.com/p1xray/pxr-url-shortener/internal/lib/short-code-generator"
)

var (
	ErrGenerateShortCode = errors.New("error generating short code")
)

// ShortURL is a business entity for a short URL.
type ShortURL struct {
	length    int
	LongUrl   string
	ShortCode string
}

// New creates a new instance of business entity for a short URL.
func New(longURL string, length int) (ShortURL, error) {
	generator := shortcodegenerator.New(length)
	shortCode, err := generator.Generate()
	if err != nil {
		return ShortURL{}, fmt.Errorf("%w: %w", ErrGenerateShortCode, err)
	}

	return ShortURL{
		length:    length,
		LongUrl:   longURL,
		ShortCode: shortCode,
	}, nil
}

// RegenerateShortCode regenerate a short code.
func (su *ShortURL) RegenerateShortCode() error {
	generator := shortcodegenerator.New(su.length)
	shortCode, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrGenerateShortCode, err)
	}

	su.ShortCode = shortCode
	return nil
}

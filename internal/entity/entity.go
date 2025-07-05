package entity

import (
	"errors"
	"fmt"
	shortcodegenerator "github.com/p1xray/pxr-url-shortener/internal/lib/short-code-generator"
	"net/url"
	"strings"
)

var (
	ErrGenerateShortCode = errors.New("error generating short code")
)

// ShortURL is a business entity for a short URL.
type ShortURL struct {
	host            string
	shortCodeLength int
	LongUrl         string
	ShortCode       string
	ShortURL        string
}

// New creates a new instance of business entity for a short URL.
func New(longURL, host string, shortCodeLength int) (ShortURL, error) {
	generator := shortcodegenerator.New(shortCodeLength)
	shortCode, err := generator.Generate()
	if err != nil {
		return ShortURL{}, fmt.Errorf("%w: %w", ErrGenerateShortCode, err)
	}

	return ShortURL{
		host:            host,
		shortCodeLength: shortCodeLength,
		LongUrl:         longURL,
		ShortCode:       shortCode,
		ShortURL:        generateShortURL(host, shortCode),
	}, nil
}

// NewWithExistingShortCode creates a new instance of business entity for a short URL with existing short code.
func NewWithExistingShortCode(longURL, shortCode, host string) ShortURL {
	return ShortURL{
		host:            host,
		shortCodeLength: len(shortCode),
		LongUrl:         longURL,
		ShortCode:       shortCode,
		ShortURL:        generateShortURL(host, shortCode),
	}
}

// RegenerateShortCode regenerate a short code.
func (su *ShortURL) RegenerateShortCode() error {
	generator := shortcodegenerator.New(su.shortCodeLength)
	shortCode, err := generator.Generate()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrGenerateShortCode, err)
	}

	su.ShortCode = shortCode
	su.ShortURL = generateShortURL(su.host, su.ShortCode)
	return nil
}

func generateShortURL(httpServerAddr, shortCode string) string {
	shortURL := url.URL{
		Host: httpServerAddr,
		Path: shortCode,
	}
	shortURLStr := strings.Trim(shortURL.String(), "//")

	return shortURLStr
}

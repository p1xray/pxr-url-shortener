package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/p1xray/pxr-url-shortener/internal/config"
	"github.com/p1xray/pxr-url-shortener/internal/dto"
	"github.com/p1xray/pxr-url-shortener/internal/entity"
	"github.com/p1xray/pxr-url-shortener/internal/storage"
	"github.com/p1xray/pxr-url-shortener/internal/storage/domain"
)

// Storage represents a URL storage.
type Storage interface {
	// URLByLongURL returns URL entity by long URL from storage.
	URLByLongURL(ctx context.Context, longURL string) (domain.URL, error)

	// IsURLExistByShortCode checks for the presence of a URL entity by short code in storage.
	IsURLExistByShortCode(ctx context.Context, shortCode string) (bool, error)

	// URLByShortCode returns URL entity by short code from storage.
	URLByShortCode(ctx context.Context, shortCode string) (domain.URL, error)

	// CreateURL creating the new URL in storage.
	CreateURL(ctx context.Context, url entity.ShortURL) error
}

// Service is a service for working with short URL.
type Service struct {
	cfg     config.ShortCodeGeneratorConfig
	storage Storage
}

// New creates a new instance of service for working with short URL.
func New(cfg config.ShortCodeGeneratorConfig, storage Storage) *Service {
	return &Service{
		cfg:     cfg,
		storage: storage,
	}
}

// Shorten generates a new unique short code for URL.
func (s *Service) Shorten(ctx context.Context, longURL, host string) (dto.Shorten, error) {
	const op = "service.Shorten"

	// first, check if a short URL exists.
	existingURL, err := s.storage.URLByLongURL(ctx, longURL)
	if err != nil && !errors.Is(err, storage.ErrEntityNotFound) {
		return dto.Shorten{}, fmt.Errorf("%s: %w", op, err)
	}

	if existingURL.ShortCode != "" {
		shortURL := entity.NewWithExistingShortCode(existingURL.LongUrl, existingURL.ShortCode, host)

		return dto.Shorten{ShortCode: shortURL.ShortCode, ShortURL: shortURL.ShortURL}, nil
	}

	// create a new short URL.
	shortURL, err := entity.New(longURL, host, s.cfg.Length)
	if err != nil {
		return dto.Shorten{}, fmt.Errorf("%s: %w", op, err)
	}

	// verify short code uniqueness.
	shortURL, err = s.verifyShortCodeUniqueness(ctx, shortURL)
	if err != nil {
		return dto.Shorten{}, fmt.Errorf("%s: %w", op, err)
	}

	// save a new short URL to storage.
	if err = s.storage.CreateURL(ctx, shortURL); err != nil {
		return dto.Shorten{}, fmt.Errorf("%s: %w", op, err)
	}

	return dto.Shorten{ShortCode: shortURL.ShortCode, ShortURL: shortURL.ShortURL}, nil
}

// LongURL returns a long URL by short code.
func (s *Service) LongURL(ctx context.Context, shortCode string) (string, error) {
	const op = "service.LongURL"

	// get URL by short code.
	url, err := s.storage.URLByShortCode(ctx, shortCode)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return url.LongUrl, nil
}

func (s *Service) verifyShortCodeUniqueness(ctx context.Context, shortURL entity.ShortURL) (entity.ShortURL, error) {
	const op = "service.verifyShortCodeUniqueness"

	isURLExist, err := s.storage.IsURLExistByShortCode(ctx, shortURL.ShortCode)
	if err != nil && !errors.Is(err, storage.ErrEntityNotFound) {
		return entity.ShortURL{}, fmt.Errorf("%s: %w", op, err)
	}

	if !isURLExist {
		return shortURL, nil
	}

	if err = shortURL.RegenerateShortCode(); err != nil {
		return entity.ShortURL{}, fmt.Errorf("%s: %w", op, err)
	}

	return s.verifyShortCodeUniqueness(ctx, shortURL)
}

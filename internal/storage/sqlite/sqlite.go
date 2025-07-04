package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"github.com/p1xray/pxr-url-shortener/internal/entity"
	"github.com/p1xray/pxr-url-shortener/internal/storage"
	"github.com/p1xray/pxr-url-shortener/internal/storage/domain"
	"time"
)

// Storage represents a URL SQLite storage.
type Storage struct {
	db *sql.DB
}

// New creates a new SQLite storage.
func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

// URLByLongURL returns URL entity by long URL from storage.
func (s *Storage) URLByLongURL(ctx context.Context, longURL string) (domain.URL, error) {
	const op = "storage.URLByLongURL"

	stmt, err := s.db.PrepareContext(ctx,
		`select
    	u.id,
    	u.long_url,
    	u.short_code,
    	u.created_at,
    	u.updated_at
		from urls u
		where u.long_url = ?;`)

	if err != nil {
		return domain.URL{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, longURL)

	var url domain.URL
	err = row.Scan(
		&url.ID,
		&url.LongUrl,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.URL{}, fmt.Errorf("%s: %w", op, storage.ErrEntityNotFound)
		}

		return domain.URL{}, fmt.Errorf("%s: %w", op, err)
	}

	return url, nil
}

// IsURLExistByShortCode checks for the presence of a URL entity by short code in storage.
func (s *Storage) IsURLExistByShortCode(ctx context.Context, shortCode string) (bool, error) {
	const op = "storage.sqlite.IsURLExistByShortCode"

	stmt, err := s.db.PrepareContext(ctx, `select * from urls u where u.short_code = ?;`)

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, shortCode)
	err = row.Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrEntityNotFound)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return true, nil
}

// URLByShortCode returns URL entity by short code from storage.
func (s *Storage) URLByShortCode(ctx context.Context, shortCode string) (domain.URL, error) {
	const op = "storage.sqlite.URLByShortCode"

	stmt, err := s.db.PrepareContext(ctx,
		`select
    	u.id,
    	u.long_url,
    	u.short_code,
    	u.created_at,
    	u.updated_at
		from urls u
		where u.short_code = ?;`)

	if err != nil {
		return domain.URL{}, fmt.Errorf("%s: %w", op, err)
	}

	row := stmt.QueryRowContext(ctx, shortCode)

	var url domain.URL
	err = row.Scan(
		&url.ID,
		&url.LongUrl,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.URL{}, fmt.Errorf("%s: %w", op, storage.ErrEntityNotFound)
		}

		return domain.URL{}, fmt.Errorf("%s: %w", op, err)
	}

	return url, nil
}

// CreateURL creating the new URL in storage.
func (s *Storage) CreateURL(ctx context.Context, url entity.ShortURL) error {
	const op = "storage.sqlite.CreateURL"

	stmt, err := s.db.PrepareContext(ctx,
		`insert into urls (long_url, short_code, created_at, updated_at)
		values (?, ?, ?, ?);`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	now := time.Now()
	_, err = stmt.ExecContext(ctx, url.LongUrl, url.ShortCode, now, now)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s: %w", op, storage.ErrURLExist)
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

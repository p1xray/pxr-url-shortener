package storage

import "errors"

var (
	ErrEntityNotFound = errors.New("entity not found")
	ErrURLExist       = errors.New("URL with same short code already exists")
)

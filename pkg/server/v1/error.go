package v1

import "errors"

var (
	// ErrNoKey for key not found
	ErrNoKey = errors.New("No key found")
	// ErrKeyExpired for keys expired
	ErrKeyExpired = errors.New("Key expired")
)

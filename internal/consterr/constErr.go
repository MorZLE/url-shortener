package consterr

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrKeyBusy  = errors.New("key Busy")
	ErrGenShort = errors.New("failed to generate short URL")
)

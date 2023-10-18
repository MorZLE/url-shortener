package consts

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrKeyBusy      = errors.New("key Busy")
	ErrGenShort     = errors.New("failed to generate short URL")
	ErrGetURL       = errors.New("failed to get URL")
	ErrDuplicateURL = errors.New("duplicate URL")
)

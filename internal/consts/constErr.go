package consts

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("key already exists")
	ErrGenShort      = errors.New("failed to generate short URL")
	ErrGetURL        = errors.New("failed to get URL")
	ErrDuplicateURL  = errors.New("duplicate URL")
	ErrBlockURL      = errors.New("delete URL")
)

package _const

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrKeyBusy  = errors.New("key Busy")
)

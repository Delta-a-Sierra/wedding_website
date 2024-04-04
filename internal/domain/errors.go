package domain

import "errors"

var (
	ErrRegistryItemNotFound = errors.New("registry item not found")
	ErrGuestNotFound        = errors.New("guest not found")
)

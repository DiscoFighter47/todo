package data

import "errors"

var (
	// ErrUserAlreadyExists ...
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("user not found")
)

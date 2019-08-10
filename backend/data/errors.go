package data

import "errors"

var (
	// ErrUserAlreadyExists ...
	ErrUserAlreadyExists = errors.New("User already exists")
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("User not found")
)

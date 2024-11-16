package errors

import "errors"

// Auth
var (
	ErrUserExisted  = errors.New("user existed")
	ErrUserNotFound = errors.New("user not found")
)

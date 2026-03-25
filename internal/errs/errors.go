package errs

import "errors"

var (
	ErrInvalidTitle   = errors.New("title is required")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternal       = errors.New("internal server error")
)
package errors

import "errors"

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrValidationFailed    = errors.New("validation failed")
)

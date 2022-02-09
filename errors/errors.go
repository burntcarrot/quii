package errors

import "errors"

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrInternalServerError  = errors.New("internal server error")
	ErrBadRequest           = errors.New("bad request")
	ErrValidationFailed     = errors.New("validation failed")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrProjectAlreadyExists = errors.New("project already exists")
	ErrTaskAlreadyExists    = errors.New("task already exists")
)

package errs

import "errors"

var (
	ErrUserDoesNotExist   = errors.New("user with provided params does not exist")
	ErrNoParams           = errors.New("search params was not provided")
	ErrFailedToCreateUser = errors.New("failed to create user")
)

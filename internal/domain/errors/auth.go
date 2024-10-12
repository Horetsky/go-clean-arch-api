package errs

import "errors"

var (
	ErrUnauthorized           = errors.New("unauthorized")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrInvalidPassword        = errors.New("invalid password")
	ErrFailedToCreateSession  = errors.New("failed to create session")
	ErrFailedToParseJWTClaims = errors.New("failed to pars jwt claims")
)

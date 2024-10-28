package errs

import "errors"

var (
	ErrUnauthorized           = errors.New("unauthorized")
	ErrFailedToVerifyEmail    = errors.New("failed to verify email")
	ErrUserAlreadyExists      = errors.New("user already exists")
	ErrInvalidPassword        = errors.New("invalid password")
	ErrFailedToDeleteAccount  = errors.New("failed to delete account")
	ErrFailedToCreateSession  = errors.New("failed to create session")
	ErrFailedToParseJWTClaims = errors.New("failed to pars jwt claims")
)

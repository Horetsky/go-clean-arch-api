package errs

import "errors"

var (
	ErrTalentAlreadyExists  = errors.New("talent already exists")
	ErrFailedToCreateTalent = errors.New("failed to create talent")
)

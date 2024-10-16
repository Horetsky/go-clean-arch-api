package errs

import "errors"

var (
	ErrRecruiterAlreadyExists  = errors.New("recruiter already exists")
	ErrFailedToCreateRecruiter = errors.New("failed to create recruiter")
)

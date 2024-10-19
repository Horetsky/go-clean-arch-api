package errs

import "errors"

var (
	ErrFailedToPostJob = errors.New("failed to post a job")
)

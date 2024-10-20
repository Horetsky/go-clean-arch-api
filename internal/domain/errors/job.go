package errs

import "errors"

var (
	ErrFailedToSendApplicationEmail = errors.New("failed to send application email")
	ErrFailedToPostJob              = errors.New("failed to post a job")
	ErrFailedToApplyJob             = errors.New("failed to apply job")
	ErrApplicationAlreadyExists     = errors.New("you have already applied for this job")
)

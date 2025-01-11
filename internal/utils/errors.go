package utils

import "errors"

var (
	ErrMapContext       error = errors.New("no Map defined in context")
	ErrSchedulerContext error = errors.New("no Scheduler defined in context")
	ErrObserverContext  error = errors.New("no Observer defined in context")
)

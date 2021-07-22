package helpers

import "errors"

var (
	// ErrInternalServerError is message for internal server error
	ErrInternalServerError = errors.New("ServerError / Internal Server Error")
	// ErrNotFound is message for not found item
	ErrNotFound = errors.New("Not found")
	// ErrConflict is message for conflict item
	ErrConflict = errors.New("Conflict")
	// ErrBadRequest is message for not bad request
	ErrBadRequest = errors.New("Bad Request")
)

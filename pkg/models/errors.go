package models

import "errors"

var (
	// ErrBadRequest will throw if the given request-body or params is not valid
	ErrBadRequest = errors.New("bad request")
	// ErrCannotBeDeleted will throw if the requested item can not be deleted
	ErrCannotBeDeleted = errors.New("cannot be deleted")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("item already exists")
	// ErrDatabaseError will throw if any database error happen
	ErrDatabaseError = errors.New("database error")
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if the requested item does not exists
	ErrNotFound = errors.New("requested item is not found")
	// ErrTimeout will throw when the request timeout
	ErrTimeout = errors.New("timeout")
)

package errors

import "errors"

var (
	ErrUsernameNotFound      = errors.New("username not found")
	ErrEmailNotFound         = errors.New("email not found")
	ErrUsernameAlreadyExist  = errors.New("username already exist")
	ErrEmailAlreadyExist     = errors.New("email already exist")
	ErrUserFirstNameRequired = errors.New("first name required")
	ErrUserLastNameRequired  = errors.New("last name required")
	ErrUserUsernameRequired  = errors.New("username required")
	ErrUserEmailRequired     = errors.New("email required")
	ErrPasswordRequired      = errors.New("password required")
)

package errors

import "errors"

var (
	ErrNoAuthorizationHeader   = errors.New("authorization header required")
	ErrInvalidTokenFormat      = errors.New("invalid token format")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrTokenExpired            = errors.New("expired token")
	ErrAccessDenied            = errors.New("access denied")
)

package domain

import "errors"

var (
	// User errors
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentionals")

	// Auth errors
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
	ErrUnauthorized = errors.New("unauthorized")

	// Room errors
	ErrRoomNotFound     = errors.New("room not found")
	ErrNotRoomMember    = errors.New("not a room member")
	ErrAlreadyMember    = errors.New("already a member")
	ErrPermissionDenied = errors.New("permission denied")

	// Message errors
	ErrMessageNotFound = errors.New("message not found")
	ErrInvalidMessage  = errors.New("invalid message")

	// Common errors
	ErrInternal     = errors.New("internal error")
	ErrInvalidInput = errors.New("invalid input")
)

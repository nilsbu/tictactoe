package util

// ErrorAnticipation is a type that is used in test tables to describe if an
// error is expected by a function when no other information about that error is
// described.
type ErrorAnticipation int

// AnyError and NoError are placeholders for an indescript error and no error.
// They are meant to be used in test tables.
const (
	AnyError ErrorAnticipation = iota
	NoError
)

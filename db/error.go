package db

import "fmt"

// An Error implements the error interface and contains data about a database
// error.
type Error struct {
	Type    ErrorType
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

func newError(t ErrorType, msg string, args ...interface{}) error {
	return &Error{
		Type:    t,
		Message: fmt.Sprintf(msg, args...),
	}
}

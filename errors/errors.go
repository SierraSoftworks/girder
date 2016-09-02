package errors

import (
	"fmt"
	"runtime/debug"
)

// Error represents an API level error object built up of a code, name and message
type Error struct {
	Code    int    `json:"code"`
	Name    string `json:"error"`
	Message string `json:"message"`
	Stack   string `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Code, e.Name, e.Message)
}

// NewError creates a new API level error object with the given properties
func NewError(code int, name, message string) *Error {
	return &Error{
		Code:    code,
		Name:    name,
		Message: message,
		Stack:   string(debug.Stack()),
	}
}

// From creates a new API level error object for a general error object
func From(err error) *Error {
	switch err.(type) {
	case *Error:
		return err.(*Error)
	}

	return Formatter.Format(err)
}

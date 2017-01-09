package errors

import "fmt"

// Error represents an API level error object built up of a code, name and message
type Error struct {
	Code       int         `json:"code"`
	Name       string      `json:"error"`
	Message    string      `json:"message"`
	Stacktrace interface{} `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.Code, e.Name, e.Message)
}

// WithCode allows you to create a new error object which inherits the properties
// of its source, with a different error code.
func (e *Error) WithCode(code int) *Error {
	return &Error{
		Code:       code,
		Name:       e.Name,
		Message:    e.Message,
		Stacktrace: e.Stacktrace,
	}
}

// WithName allows you to create a new error object which inherits the properties
// of its source, with a different error name.
func (e *Error) WithName(name string) *Error {
	return &Error{
		Code:       e.Code,
		Name:       name,
		Message:    e.Message,
		Stacktrace: e.Stacktrace,
	}
}

// WithMessage allows you to create a new error object which inherits the properties
// of its source, with a different error message.
func (e *Error) WithMessage(message string) *Error {
	return &Error{
		Code:       e.Code,
		Name:       e.Name,
		Message:    message,
		Stacktrace: e.Stacktrace,
	}
}

// NewError creates a new API level error object with the given properties
func NewError(code int, name, message string) *Error {
	return &Error{
		Code:       code,
		Name:       name,
		Message:    message,
		Stacktrace: Stacktrace.Get(),
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

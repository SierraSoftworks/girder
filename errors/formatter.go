package errors

type ErrorFormatter interface {
	Format(err error) *Error
}

var Formatter ErrorFormatter = &DefaultFormatter{}

type DefaultFormatter struct {
}

func (f *DefaultFormatter) Format(err error) *Error {
	return &Error{
		Code:    500,
		Name:    "Server Error",
		Message: "An unexpected error occured while processing your request. Please check it and try again.",
	}
}

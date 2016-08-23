package errors

// ServerError creates a 500 equivalent error object
func ServerError() *Error {
	return &Error{
		Code:    500,
		Name:    "Server Error",
		Message: "We were unable to process your request because of an unexpected error, please check it and try again.",
	}
}

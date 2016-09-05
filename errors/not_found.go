package errors

// NotFound creates a 404 equivalent error object
func NotFound() *Error {
	return NewError(
		404,
		"Not Found",
		"The entity you were looking for could not be found, please check your request and try again.",
	)
}

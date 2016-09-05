package errors

// BadRequest creates a 400 equivalent error object
func BadRequest() *Error {
	return NewError(
		400,
		"Bad Request",
		"The request you made was not correctly formatted, please check it and try again.",
	)
}

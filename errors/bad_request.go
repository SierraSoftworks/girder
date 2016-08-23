package errors

// BadRequest creates a 400 equivalent error object
func BadRequest() *Error {
	return &Error{
		Code:    400,
		Name:    "Bad Request",
		Message: "The request you made was not correctly formatted, please check it and try again.",
	}
}

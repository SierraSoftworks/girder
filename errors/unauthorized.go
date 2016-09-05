package errors

// Unauthorized creates a 401 equivalent error object
func Unauthorized() *Error {
	return NewError(
		401,
		"Unauthorized",
		"You have failed to provide a valid authentication token with your request.",
	)
}

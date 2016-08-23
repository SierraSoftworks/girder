package errors

// Unauthorized creates a 401 equivalent error object
func Unauthorized() *Error {
	return &Error{
		Code:    401,
		Name:    "Unauthorized",
		Message: "You have failed to provide a valid authentication token with your request.",
	}
}

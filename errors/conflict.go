package errors

// Conflict creates a 409 equivalent error object
func Conflict() *Error {
	return &Error{
		Code:    409,
		Name:    "Conflict",
		Message: "You attempted to create something with an identifier which already exists.",
	}
}

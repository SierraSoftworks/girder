package errors

// Conflict creates a 409 equivalent error object
func Conflict() *Error {
	return NewError(
		409,
		"Conflict",
		"You attempted to create something with an identifier which already exists.",
	)
}

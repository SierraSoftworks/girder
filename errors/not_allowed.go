package errors

// NotAllowed creates a 403 equivalent error object
func NotAllowed() *Error {
	return NewError(
		403,
		"Not Allowed",
		"You do not have permission to access the entity you requested.",
	)
}

package errors

// NotAllowed creates a 403 equivalent error object
func NotAllowed() *Error {
	return &Error{
		Code:    403,
		Name:    "Not Allowed",
		Message: "You do not have permission to access the entity you requested.",
	}
}

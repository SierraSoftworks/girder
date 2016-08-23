package girder

// User represents an authenticated user
type User interface {
	GetID() string
	GetPermissions() []string
}

package girder

import (
	"net/http"

	"github.com/SierraSoftworks/gatekeeper"
	"github.com/SierraSoftworks/girder/errors"
)

// Context represents an API request's context
type Context struct {
	Request         *http.Request
	Vars            map[string]string
	ResponseHeaders http.Header
	StatusCode      int
	User            User
	Permissions     *gatekeeper.Matcher

	response http.ResponseWriter
}

// ReadBody will deserialize the request's body into the given object
func (c *Context) ReadBody(into interface{}) error {
	err := parseJSON(into, c)
	if err != nil {
		return errors.BadRequest()
	}

	return nil
}

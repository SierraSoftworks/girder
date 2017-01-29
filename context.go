package girder

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"

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
	Formatter       Formatter
	Parser          Parser

	response http.ResponseWriter
}

// ReadBody will deserialize the request's body into the given object
func (c *Context) ReadBody(into interface{}) error {
	if c.Parser == nil {
		err := fmt.Errorf("no parser set")

		log.
			WithError(err).
			WithField("request", c.Request).
			Error("No parser available for this context")

		return err
	}

	err := c.Parser.Read(into, c.Request.Body)
	if err != nil {
		return errors.BadRequest()
	}

	return nil
}

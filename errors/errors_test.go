package errors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestErrorInterface(c *C) {
	err := NewError(400, "test", "test message")
	var e error
	e = err

	c.Check(e.Error(), Equals, "400 test: test message")
}

func (s *TestSuite) TestNewError(c *C) {
	err := NewError(400, "Bad Request", "Test message")
	c.Assert(err, NotNil)
	c.Check(err.Code, Equals, 400)
	c.Check(err.Name, Equals, "Bad Request")
	c.Check(err.Message, Equals, "Test message")
}

func (s *TestSuite) TestErrorSerialization(c *C) {
	b := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(b)

	err := NewError(400, "test", "test message")

	c.Assert(enc.Encode(err), IsNil)

	c.Assert(strings.TrimSpace(b.String()), Equals, `{"code":400,"error":"test","message":"test message"}`)
}

func (s *TestSuite) TestFrom(c *C) {
	err := From(fmt.Errorf("test error"))
	c.Assert(err, NotNil)
	c.Check(err.Code, Equals, 500)
	c.Check(err.Name, Equals, "Server Error")
	c.Check(err.Message, Equals, "An unexpected error occured while processing your request. Please check it and try again.")

	err = From(NewError(400, "test", "test message"))
	c.Assert(err, NotNil)
	c.Check(err.Code, Equals, 400)
	c.Check(err.Name, Equals, "test")
	c.Check(err.Message, Equals, "test message")
}

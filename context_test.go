package girder

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/SierraSoftworks/girder/errors"
	. "gopkg.in/check.v1"
)

type TestReadBodyData struct {
	Test int `json:"test"`
}

func (s *TestSuite) TestReadBody(c *C) {
	ctx := &Context{
		Request: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"test":1}`)),
		},
	}

	var data TestReadBodyData
	err := ctx.ReadBody(&data)
	c.Assert(err, IsNil)
	c.Check(data, DeepEquals, TestReadBodyData{Test: 1})
}

func (s *TestSuite) TestReadBodyMalformedData(c *C) {
	ctx := &Context{
		Request: &http.Request{
			Body: ioutil.NopCloser(bytes.NewBufferString(`{test:1}`)),
		},
	}

	var data TestReadBodyData
	err := ctx.ReadBody(&data)
	c.Assert(err, NotNil)

	e := errors.From(err)
	c.Check(e.Code, Equals, 400)
	c.Check(e.Name, Equals, "Bad Request")
}

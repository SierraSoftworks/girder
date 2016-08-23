package girder

import (
	"net/http"

	"github.com/SierraSoftworks/gatekeeper"
	"github.com/SierraSoftworks/girder/errors"
	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestIsAuthenticated(c *C) {
	ctx := &Context{}
	c.Check(ctx.IsAuthenticated(), Equals, false)

	ctx = &Context{
		Permissions: gatekeeper.NewMatcher(),
		User: &testUser{
			id:          "bob",
			permissions: []string{"x"},
		},
	}

	c.Check(ctx.IsAuthenticated(), Equals, true)
}

func (s *TestSuite) TestRequireAuthentication(c *C) {
	ctx := &Context{
		Permissions: gatekeeper.NewMatcher(),
		Request: &http.Request{
			Header: http.Header{},
		},
	}

	h := NewHandler(nil)
	h.RequireAuthentication(func(token *AuthorizationToken) (User, error) {
		return &testUser{
			id:          token.Value,
			permissions: []string{},
		}, nil
	})

	c.Check(h.Preprocessors, HasLen, 1)

	proc := h.Preprocessors[0]
	err := proc(ctx)
	c.Assert(err, NotNil)
	e := errors.From(err)
	c.Check(e.Code, Equals, 401)
	c.Check(e.Name, Equals, "Unauthorized")

	ctx.Request.Header.Set("Authorization", "Token test")
	err = proc(ctx)
	c.Assert(err, IsNil)
	c.Check(ctx.User, DeepEquals, &testUser{id: "test", permissions: []string{}})
}

package girder

import (
	"net/http"

	"github.com/SierraSoftworks/girder/errors"
	. "gopkg.in/check.v1"
)

type TestUser struct {
	id string
}

func (u *TestUser) GetID() string {
	return u.id
}

func (u *TestUser) HasPermission(permission string) bool {
	return true
}

func (s *TestSuite) TestIsAuthenticated(c *C) {
	ctx := &Context{}
	c.Check(ctx.IsAuthenticated(), Equals, false)

	ctx = &Context{
		User: &TestUser{
			id: "bob",
		},
	}

	c.Check(ctx.IsAuthenticated(), Equals, true)
}

func (s *TestSuite) TestRequireAuthentication(c *C) {
	ctx := &Context{
		Request: &http.Request{
			Header: http.Header{},
		},
	}

	h := NewHandler(nil)
	h.RequireAuthentication(func(token *AuthorizationToken) (User, error) {
		return &TestUser{
			id: token.Value,
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
	c.Check(ctx.User, DeepEquals, &TestUser{id: "test"})
}

package girder

import "github.com/SierraSoftworks/girder/errors"

// IsAuthenticated asserts that the request has been authenticated
func (c *Context) IsAuthenticated() bool {
	return c.User != nil
}

// RequireAuthentication configures this handler to require authentication
func (h *Handler) RequireAuthentication(getUser func(token *AuthorizationToken) (User, error)) *Handler {
	return h.RegisterPreprocessors(func(c *Context) error {
		token := c.GetAuthToken()
		if token != nil {
			user, err := getUser(token)
			if err != nil {
				return err
			}

			c.User = user
		}

		if !c.IsAuthenticated() {
			return errors.Unauthorized()
		}

		return nil
	})
}

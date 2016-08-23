package girder

import "github.com/SierraSoftworks/girder/errors"

// RequirePermission configures this handler to require the provided permission for all requests
func (h *Handler) RequirePermission(permission string) *Handler {
	return h.RegisterPreprocessors(func(c *Context) error {
		if c.User == nil {
			return errors.Unauthorized()
		}

		if !c.User.HasPermission(permission) {
			return errors.NotAllowed()
		}

		return nil
	})
}

package girder

import "github.com/SierraSoftworks/girder/errors"

// RequirePermission configures this handler to require the provided permission for all requests
func (h *Handler) RequirePermission(permissions ...string) *Handler {
	return h.RegisterPreprocessors(func(c *Context) error {
		if !c.Permissions.CanAll(permissions...) {
			return errors.NotAllowed()
		}

		return nil
	})
}

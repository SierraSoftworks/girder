package girder

import log "github.com/sirupsen/logrus"

// LogRequests adds debug logging on all requests to simplify diagnosing problems
func (h *Handler) LogRequests() *Handler {
	return h.RegisterPreprocessors(func(c *Context) error {
		log.WithFields(log.Fields{
			"headers": c.Request.Header,
			"vars":    c.Vars,
		}).Infof("%s %s", c.Request.Method, c.Request.URL.String())

		return nil
	})
}

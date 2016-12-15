package girder

import (
	"net/http"

	"github.com/SierraSoftworks/gatekeeper"
	"github.com/SierraSoftworks/girder/errors"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// HandlerFunc represents a function that handles an API request and returns a result
type HandlerFunc func(c *Context) (interface{}, error)

// Handler represents an API request handler
type Handler struct {
	HandleFunc HandlerFunc

	Preprocessors []Preprocessor
}

// NewHandler creates a new HTTP compatible handler for the given API HandlerFunc
func NewHandler(handle HandlerFunc) *Handler {
	return &Handler{
		HandleFunc:    handle,
		Preprocessors: []Preprocessor{},
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		Request:         r,
		ResponseHeaders: w.Header(),
		Vars:            mux.Vars(r),
		StatusCode:      200,
		Permissions:     gatekeeper.NewMatcher().WithContext(mux.Vars(r)),
		Formatter:       &JSONFormatter{},

		response: w,
	}

	c.ResponseHeaders.Set("Content-Type", "application/json; charset=utf-8")

	for _, preprocessor := range h.Preprocessors {
		if err := preprocessor(c); err != nil {
			e := errors.From(err)
			w.WriteHeader(e.Code)
			if err := writeJSON(e, c); err != nil {
				log.Error("Failed to encode error to JSON", err)
			}

			return
		}
	}

	res, err := h.HandleFunc(c)

	if err != nil {
		e := errors.From(err)
		w.WriteHeader(e.Code)
		if err := writeJSON(e, c); err != nil {
			log.Error("Failed to encode error to JSON", err)
		}

		return
	}

	w.WriteHeader(c.StatusCode)
	if res != nil {
		if c.Formatter == nil {
			log.
				WithField("response", res).
				Error("No formatter available for this context", err)
		} else if err := c.Formatter.Write(res, c.response); err != nil {
			log.
				WithField("response", res).
				WithField("formatter", c.Formatter).
				Error("Failed to encode response", err)
		}
	}
}

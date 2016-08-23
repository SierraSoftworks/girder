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

		response: w,
	}

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
		if err := writeJSON(res, c); err != nil {
			log.WithField("response", res).Error("Failed to encode response to JSON", err)
		}
	}
}

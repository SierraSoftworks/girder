package girder

// Preprocessor entries are executed during the request's pipeline and
// can pre-emptively kick out of the processing stage by raising
// an error.
type Preprocessor func(c *Context) error

// RegisterPreprocessors will register a new preprocessor on this handler's
// preprocessor pipeline.
func (h *Handler) RegisterPreprocessors(preprocessors ...Preprocessor) *Handler {
	h.Preprocessors = append(h.Preprocessors, preprocessors...)
	return h
}

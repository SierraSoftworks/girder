package girder

import "io"

// Formatter is responsible for formatting the data that a handler
// returns.
type Formatter interface {
	Write(data interface{}, into io.Writer) error
}

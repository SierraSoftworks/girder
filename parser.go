package girder

import "io"

// Parser is responsible for parsing the data that is received by
// a handler.
type Parser interface {
	Read(data interface{}, from io.Reader) error
}

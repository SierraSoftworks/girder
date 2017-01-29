package girder

import (
	"encoding/json"
	"io"
)

func writeJSON(data interface{}, c *Context) error {
	enc := json.NewEncoder(c.response)
	return enc.Encode(data)
}

// JSONFormatter is a formatter implementation which encodes the
// response data using the JSON format.
type JSONFormatter struct{}

func (f *JSONFormatter) Write(data interface{}, into io.Writer) error {
	return json.NewEncoder(into).Encode(data)
}

func (f *JSONFormatter) Read(data interface{}, from io.Reader) error {
	return json.NewDecoder(from).Decode(data)
}

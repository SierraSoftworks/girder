package girder

import (
	"encoding/json"
	"io"
)

func parseJSON(data interface{}, c *Context) error {
	dec := json.NewDecoder(c.Request.Body)
	return dec.Decode(data)
}

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

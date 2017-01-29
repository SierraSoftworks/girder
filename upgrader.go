package girder

import "net/http"

// Upgrader represents a feature which will upgrade an HTTP
// connection to another protocol such as WebSocket.
type Upgrader interface {
	Upgrade(c *Context, w http.ResponseWriter) error
}

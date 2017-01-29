package girder

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketUpgrader struct {
	headers  http.Header
	upgrader websocket.Upgrader
	handler  func(conn *websocket.Conn)
}

func (u *WebSocketUpgrader) Upgrade(c *Context, w http.ResponseWriter) error {
	conn, err := u.upgrader.Upgrade(w, c.Request, u.headers)
	if err != nil {
		return err
	}

	go func() {
		u.handler(conn)
	}()

	return nil
}

func (c *Context) WebSocket(headers http.Header, handler func(conn *websocket.Conn)) (Upgrader, error) {
	return &WebSocketUpgrader{
		headers:  headers,
		upgrader: websocket.Upgrader{},
		handler:  handler,
	}, nil
}

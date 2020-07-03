package progolog

import (
	"net/url"

	"github.com/gorilla/websocket"
)

// LogsSocket describes a WebSocket connection to the logs service.
type LogsSocket interface {
}

// Returns a websocket connection to the loom service.
func newLoomConnection() *websocket.Conn {
	host := "" //config.Get("LOOM_HOST") + ":" + config.Get("LOOM_PORT")
	u := url.URL{Scheme: "ws", Host: host, Path: "/store"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	return c
}

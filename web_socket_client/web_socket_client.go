package web_socket_client

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocket struct {
	Connection *websocket.Conn
}

func (w *WebSocket) Connect(address string, header http.Header) (err error) {
	conn, _, err := websocket.DefaultDialer.Dial(address, header)
	if err != nil {
		return err
	}
	w.Connection = conn

	return err
}

func (w *WebSocket) Close() {
	w.Connection.Close()
}

func (w *WebSocket) SendMessage(message string) error {
	return w.Connection.WriteMessage(websocket.TextMessage, []byte(message))
}

func (w *WebSocket) ReadOneMessage() (string, error) {
	_, data, err := w.Connection.ReadMessage()
	return string(data), err
}

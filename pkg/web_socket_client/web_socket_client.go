package web_socket_client

import (
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocket struct {
	Connection *websocket.Conn
	messages   chan *string
}

func (w *WebSocket) startRead() {
	for {
		_, payload, err := w.Connection.ReadMessage()
		if err != nil {
			break
		}
		message := string(payload)
		w.messages <- &message
	}

	close(w.messages)
}

func (w *WebSocket) Connect(address string, header http.Header) (err error) {
	conn, _, err := websocket.DefaultDialer.Dial(address, header)
	if err != nil {
		return err
	}
	w.Connection = conn
	w.messages = make(chan *string, 100)
	go w.startRead()

	return err
}

func (w *WebSocket) Close() {
	_ = w.Connection.Close()
}

func (w *WebSocket) SendMessage(message string) error {
	return w.Connection.WriteMessage(websocket.TextMessage, []byte(message))
}

func (w *WebSocket) ReadOneMessage() (*string, error) {
	message := <-w.messages
	if message == nil {
		return nil, errors.New("connection is closed")
	}

	return message, nil
}

func (w *WebSocket) Messages() (chan *string, error) {
	if w.messages == nil {
		return nil, errors.New("connection is not established")
	}

	return w.messages, nil
}

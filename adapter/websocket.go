package adapter

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	. "github.com/imdevin567/jargon/protocol"
)

// WSAdapter ...
type WSAdapter struct {
	*Adapter
	Path      string
	localChan chan []byte
}

// NewWSAdapter ...
func NewWSAdapter(direction Direction, host string, port int, path string) *WSAdapter {
	adapter := NewAdapter(direction, host, port, WS)
	return &WSAdapter{
		Adapter: adapter,
		Path:    path,
	}
}

// Start ...
func (ws *WSAdapter) Start() {
	fmt.Println("No setup required for WS adapter")
}

// Input = create server ...
func (ws *WSAdapter) Input(c chan []byte) {
	ws.createServer(c)
}

// Output = connect to server ...
func (ws *WSAdapter) Output(c chan []byte) {
	ws.createServer(c)
}

// createServer ...
func (ws *WSAdapter) createServer(c chan []byte) {
	ws.localChan = c
	mux := http.NewServeMux()
	mux.HandleFunc(ws.Path, ws.handleConn)
	http.ListenAndServe(fmt.Sprintf("%v:%v%v", ws.Host, ws.Port, ws.Path), mux)
}

// handleConn
func (ws *WSAdapter) handleConn(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		// TODO: Handle error
	}

	defer conn.Close()

	if ws.Direction == Input {
		ws.readWSMessages(conn)
	} else if ws.Direction == Output {
		ws.writeWSMessages(conn)
	} else {
		fmt.Println("What direction are you trying to do here...")
	}
}

// readWSMessages ...
func (ws *WSAdapter) readWSMessages(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// TODO: Handle error
		}

		ws.localChan <- msg
	}
}

// writeWSMessages ...
func (ws *WSAdapter) writeWSMessages(conn *websocket.Conn) {
	for {
		msg := <-ws.localChan
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			// TODO: Handle error
		}
	}
}

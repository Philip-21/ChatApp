package websockets

import (
	"chatapp/internal/config"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

type Repo struct {
	app *config.AppConfig //accessing and modifying properties for the app
}

// Upgrader specifies parameters for upgrading an HTTP connection to a WebSocket connection
// making the connection to  persist
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	//allows request to be made from react server
	//allows all connections
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (m *Repo) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	/*
			Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		      The responseHeader is included in the response to the client's upgrade reques
	*/
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		m.app.ErrorLog.Println(err)
		return nil, err
	}
	return ws, nil
}

// Reader Listens for new messages to our webSockets
func (m *Repo) Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			m.app.ErrorLog.Println(err)
			return
		}
		m.app.InfoLog.Println(string(p))
		//a helper method for getting a writer using NextWriter, writing the message and closing the writer
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			m.app.ErrorLog.Println(err)
			return
		}
	}

}

func (m *Repo) Writer(conn *websocket.Conn) {
	for {
		m.app.InfoLog.Println("Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			m.app.ErrorLog.Println(err)
			return
		}
		//returns a writer for the next message to send
		w, err := conn.NextWriter(messageType)
		if err != nil {
			m.app.ErrorLog.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			m.app.ErrorLog.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			m.app.ErrorLog.Println(err)
			return
		}
	}
}

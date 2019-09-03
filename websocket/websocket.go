package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rafikbelas/rocky/server"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return ws, err
	}

	return ws, nil
}

func updateReport(conn *websocket.Conn) {

	response, err := server.GetResponse()
	if err != nil {
		fmt.Println(err)
	}

	jsonString, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); 
	err != nil {
		fmt.Println(err)
		return
	}
}

func Writer(conn *websocket.Conn) {

	for true {
		updateReport(conn)
        time.Sleep(5 * time.Second)
    }
}


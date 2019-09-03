package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/rafikbelas/rocky/websocket"
)

func reports(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/reports", reports)
}

func main() {

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

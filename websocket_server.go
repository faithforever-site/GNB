package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)

func startWebSocketServer() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket server running on :8081")
	http.ListenAndServe(":8081", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	clients[conn] = true
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		broadcast(msg)
	}
}

func broadcast(msg []byte) {
	for client := range clients {
		client.WriteMessage(websocket.TextMessage, msg)
	}
}

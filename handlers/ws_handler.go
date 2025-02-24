package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

// Handle WebSocket connections
func ChatWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Gagal meng-upgrade ke WebSocket:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}

		broadcast <- string(msg)
	}
}

// Fungsi untuk broadcast pesan ke semua klien
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

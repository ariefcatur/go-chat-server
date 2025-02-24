package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-chat-server/config"
	"go-chat-server/handlers"
	"go-chat-server/routes"
	"log"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Koneksi ke Database
	config.ConnectDB()

	// Jalankan WebSocket Broadcaster
	go handlers.HandleMessages()

	// Jalankan Server
	r := routes.SetupRouter()
	port := ":8080"
	r.Run(port)
	go routes.OpenBrowser("http://localhost:8080/static/login.html")
	fmt.Println("http://localhost:8080/static/login.html")
	go log.Fatal(http.ListenAndServe("http://localhost:8080/static/login.html", nil))

}

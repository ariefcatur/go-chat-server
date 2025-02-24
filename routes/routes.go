package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chat-server/handlers"
	"go-chat-server/middleware"
	"os/exec"
	"runtime"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Sajikan file statis dari folder frontend
	r.Static("/static", "./frontend")

	// API Routes
	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		// Protected routes (butuh JWT)
		auth := api.Group("/")
		auth.Use(middleware.AuthMiddleware())
		auth.GET("/messages", handlers.GetMessages)
		auth.POST("/messages", handlers.SendMessage)
	}

	// Jika tidak ada route yang cocok, sajikan index.html
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	// WebSocket route
	r.GET("/ws", handlers.ChatWebSocket)

	return r
}

// Fungsi untuk membuka browser otomatis
func OpenBrowser(url string) {
	go func() {
		time.Sleep(1 * time.Second) // Tunggu 1 detik agar server siap
		var cmd *exec.Cmd

		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", url)
		case "darwin": // macOS
			cmd = exec.Command("open", url)
		case "linux":
			cmd = exec.Command("xdg-open", url)
		default:
			fmt.Println("Tidak dapat membuka browser, silakan buka secara manual:", url)
			return
		}

		// Pastikan perintah dieksekusi dengan `.Run()`
		err := cmd.Run()
		if err != nil {
			fmt.Println("Gagal membuka browser:", err)
		}
	}()
}

package main

import (
	// "log"

	"example/config"
	"example/internal/logger"
	"example/internal/server"
)

func main() {
	config.LoadEnv()
	log := logger.NewLogger()
	app := server.NewServer(log)

	log.Info("🚀 Server starting on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("❌ Server error: ", err)
	}
}

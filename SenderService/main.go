package main

import (
	"SenderService/bootstrap"
	"SenderService/server"
	"SenderService/utilities"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	utilities.Error(err, "Error loading .env file")
	// Print start message
	utilities.StylePrint("Sender", "Service")
	// Initialize application
	bootstrap.RunRedis()
	server.RunGin()

	// Wait for interrupt signal to gracefully shut down the server with
	fmt.Println("program will wait for a SIGINT (Ctrl+C)")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}

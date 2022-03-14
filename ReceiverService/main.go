package main

import (
	"ReceiverService/application"
	"ReceiverService/bootstrap"
	"ReceiverService/utilities"
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
	utilities.StylePrint("Receiver", "Service")
	// Initialize application
	bootstrap.RunRedis()
	bootstrap.RunMysql()
	// Run application
	application.StartConsumeOrder()

	// Wait for interrupt signal to gracefully shut down the server with
	fmt.Println("program will wait for a SIGINT (Ctrl+C)")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}

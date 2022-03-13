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
	err := godotenv.Load()
	utilities.Error(err, "Error loading .env file")
	utilities.StylePrint("Sender", "Service")
	bootstrap.RunRedis()
	server.RunGin()
	fmt.Println("program will wait for a SIGINT (Ctrl+C)")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}

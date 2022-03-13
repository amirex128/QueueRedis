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
	err := godotenv.Load()
	utilities.Error(err, "Error loading .env file")
	utilities.StylePrint("Receiver", "Service")
	bootstrap.RunRedis()
	bootstrap.RunMysql()
	application.ConsumeOrder()

	fmt.Println("program will wait for a SIGINT (Ctrl+C)")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
}

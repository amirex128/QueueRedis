package application

import (
	"ReceiverService/bootstrap"
	"ReceiverService/utilities"
	"fmt"
	"github.com/adjust/rmq/v4"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	prefetchLimit = 1000
	pollDuration  = 100 * time.Millisecond
)

func StartConsumeOrder() {
	// initialize the bootstrap for start consuming
	queue := bootstrap.TaskQueue
	err := queue.StartConsuming(prefetchLimit, pollDuration)
	utilities.Error(err)

	// Add function to handle messages
	_, err = queue.AddConsumerFunc(fmt.Sprintf("consume"), func(delivery rmq.Delivery) {
		CreateOrder(delivery.Payload())
	})
	utilities.Error(err)

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	exitProcess()
}

func exitProcess() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals
	go func() {
		<-signals
		os.Exit(1)
	}()
}

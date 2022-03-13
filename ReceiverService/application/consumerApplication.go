package application

import (
	"ReceiverService/bootstrap"
	"ReceiverService/model/dto"
	"ReceiverService/utilities"
	"encoding/json"
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

func ConsumeOrder() {

	queue := bootstrap.TaskQueue
	err := queue.StartConsuming(prefetchLimit, pollDuration)
	utilities.Error(err)

	_, err = queue.AddConsumerFunc(fmt.Sprintf("consume"), func(delivery rmq.Delivery) {
		order := new(dto.OrderRequest)
		if err := json.Unmarshal([]byte(delivery.Payload()), order); err != nil {
			fmt.Println("Error unmarshalling order: ", err)
		}
		CreateOrder(order)
	})
	utilities.Error(err)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	defer signal.Stop(signals)

	<-signals
	go func() {
		<-signals
		os.Exit(1)
	}()
}

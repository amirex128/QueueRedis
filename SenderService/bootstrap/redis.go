package bootstrap

import (
	"SenderService/utilities"
	"fmt"
	"github.com/adjust/rmq/v4"
	"os"
	"sync"
)

// singleton

var once sync.Once
var TaskQueue rmq.Queue

func RunRedis() {
	once.Do(func() {

		host := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
		errChan := make(chan error)
		connection, err := rmq.OpenConnection("snapp", "tcp", host, 1, errChan)
		utilities.Error(err)
		TaskQueue, err = connection.OpenQueue("orders")
		utilities.Error(err)
	})

}

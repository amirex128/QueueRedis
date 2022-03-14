package application

import (
	"SenderService/bootstrap"
	"SenderService/model/dto"
	"encoding/json"
)

func OrderToQueue(order *dto.OrderRequest) bool {

	// get redis client
	queue := bootstrap.TaskQueue
	// marshal order to json
	marshal, err := json.Marshal(*order)
	if err != nil {
		return false
	}
	// push order to queue
	err = queue.PublishBytes(marshal)
	if err != nil {
		return false
	}
	return true
}

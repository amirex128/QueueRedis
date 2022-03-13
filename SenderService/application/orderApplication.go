package application

import (
	"SenderService/bootstrap"
	"SenderService/model/dto"
	"encoding/json"
	"fmt"
)

func OrderToQueue(order *dto.OrderRequest) bool {

	queue := bootstrap.TaskQueue
	marshal, err := json.Marshal(*order)
	if err != nil {
		return false
	}
	fmt.Println(string(marshal))
	err = queue.PublishBytes(marshal)
	if err != nil {
		return false
	}
	return true
}

package application

import (
	"ReceiverService/bootstrap"
	"ReceiverService/model/dto"
	"ReceiverService/model/orm"
	"encoding/json"
	"fmt"
)

func CreateOrder(data string) bool {
	// unmarshal json data to dto
	order := unsterilized(data)
	// create order
	db := bootstrap.MysqlClient
	result := db.Create(&order)
	if result.Error != nil {
		fmt.Printf("Create order failed, error: %v", result.Error)
		return false
	}
	return true
}

func unsterilized(data string) orm.Order {
	orderRequest := new(dto.OrderRequest)
	if err := json.Unmarshal([]byte(data), orderRequest); err != nil {
		fmt.Println("Error unmarshalling order: ", err)
	}
	order := orm.Order{
		OrderId: orderRequest.OrderId,
		Price:   orderRequest.Price,
		Title:   orderRequest.Title,
	}
	return order
}

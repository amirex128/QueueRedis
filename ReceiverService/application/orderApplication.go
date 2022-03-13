package application

import (
	"ReceiverService/bootstrap"
	"ReceiverService/model/dto"
	"ReceiverService/model/orm"
	"fmt"
)

func CreateOrder(orderRequest *dto.OrderRequest) {
	db := bootstrap.MysqlClient
	order := &orm.Order{
		OrderId: orderRequest.OrderId,
		Price:   orderRequest.Price,
		Title:   orderRequest.Title,
	}
	result := db.Create(&order)
	if result.Error != nil {
		fmt.Printf("Create order failed, error: %v", result.Error)
	}

}

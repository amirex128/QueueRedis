package controller

import (
	"SenderService/application"
	"SenderService/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Order(c *gin.Context) {
	params := &dto.OrderRequest{}
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error :": err.Error()})
		return
	}
	result := application.OrderToQueue(params)
	if result {
		c.JSON(200, gin.H{
			"message": "order is created",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "order is not created",
	})
	return
}

package controller

import (
	"SenderService/application"
	"SenderService/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Order(c *gin.Context) {
	// Data model binding from OrderRequest
	params, done := getParams(c)
	if done {
		return
	}
	// Call Order sender service
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

// Get params from request
func getParams(c *gin.Context) (*dto.OrderRequest, bool) {
	params := &dto.OrderRequest{}
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error :": err.Error()})
		return nil, true
	}
	return params, false
}

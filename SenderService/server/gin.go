package server

import (
	"SenderService/controller"
	"SenderService/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func RunGin() {
	// Initialize the Gin router as the default one
	r := InitGin()

	api := r.Group("api")
	{
		api.POST("order", controller.Order)
	}

	err := r.Run(os.Getenv("SERVER_PORT"))
	utilities.Error(err)
}

func InitGin() *gin.Engine {
	gin.ForceConsoleColor()

	// Create logger for Gin
	file, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	r := gin.Default()
	r.Use(gin.Logger())

	// Set Recover middleware for preventing panic
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	return r
}

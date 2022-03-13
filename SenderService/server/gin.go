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
	gin.ForceConsoleColor()

	file, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(file)

	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	api := r.Group("api")
	{
		api.POST("order", controller.Order)
	}

	err := r.Run(os.Getenv("SERVER_PORT"))
	utilities.Error(err)
}

package application

import (
	"SenderService/bootstrap"
	"SenderService/model/dto"
	"SenderService/utilities"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderToQueue(t *testing.T) {

	got := OrderToQueue(&dto.OrderRequest{
		OrderId: 1,
		Price:   100,
		Title:   "test",
	})
	want := true
	assert.Equal(t, want, got)
}

func init() {
	err := godotenv.Load("../.env")
	utilities.Error(err, "Error loading .env file")
	bootstrap.RunRedis()
}

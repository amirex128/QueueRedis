package application

import (
	"ReceiverService/bootstrap"
	"ReceiverService/model/orm"
	"ReceiverService/utilities"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateOrder(t *testing.T) {

	got := CreateOrder("")
	want := true
	assert.Equal(t, want, got)
}

func TestUnsterilized(t *testing.T) {
	got := unsterilized("")
	want := orm.Order{
		OrderId: 1,
		Price:   1000,
		Title:   "test",
	}
	assert.Equal(t, want, got)
}
func init() {
	err := godotenv.Load("../.env")
	utilities.Error(err, "Error loading .env file")
	bootstrap.RunRedis()
	bootstrap.RunMysql()
}

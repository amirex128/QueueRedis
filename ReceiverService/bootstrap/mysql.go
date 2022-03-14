package bootstrap

import (
	"ReceiverService/model/orm"
	"ReceiverService/utilities"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"sync"
)

var mysqlOnce sync.Once
var MysqlClient *gorm.DB

func RunMysql() {
	// Create the database connection for once
	mysqlOnce.Do(func() {
		var err error
		mysqlHost := os.Getenv("MYSQL_HOST")
		mysqlUser := os.Getenv("MYSQL_USER")
		mysqlPassword := os.Getenv("MYSQL_PASSWORD")
		mysqlPort := os.Getenv("MYSQL_PORT")
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlUser, mysqlPassword, mysqlHost, mysqlPort, "order_management")
		MysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		utilities.Error(err, "Failed to connect to MySQL accountingDB")
		err = MysqlClient.AutoMigrate(&orm.Order{})
		utilities.Error(err)

	})
}

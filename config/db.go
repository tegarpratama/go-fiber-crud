package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "parseTime=true&loc=Asia%2FJakarta")

	DB, err = gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("🚀 Connected successfully to the database")
}

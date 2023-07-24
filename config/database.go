package config

import (
	"fmt"
	"simple-crud-golang/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	// fmt.Printf("123123 %s" + config.Dbhost)

	sqlInfo := fmt.Sprint("host=localhost port=5432 user=postgres password=postgres dbname=test sslmode=disable")

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Connection to DB is Successful.")

	return db
}

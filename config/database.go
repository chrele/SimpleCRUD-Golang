package config

import (
	"fmt"
	"simple-crud-golang/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config *Config) *gorm.DB {
	// fmt.Printf("123123 %s" + config.Dbhost)

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Dbhost, config.Dbport, config.Dbusername, config.Dbpassword, config.Dbname)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Connection to DB is Successful.")

	return db
}

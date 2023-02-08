package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectTODatabase() {
	var err error
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=soib dbname=interview"), &gorm.Config{})
	if err != nil {
		fmt.Println("Datatbase connection faild")
	}
	db.AutoMigrate()
}

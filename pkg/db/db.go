package db

import (
	"fmt"

	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBS *gorm.DB

func ConnectTODatabase() {
	var err error
	DBS, err = gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=soib dbname=interview"), &gorm.Config{})
	if err != nil {
		fmt.Println("Datatbase connection faild")
	}
	DBS.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Groupmember{},
	)
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `gorm:"type:varchar(50);not null;"`
	Lastname  string `gorm:"type:varchar(50);not null;"`
	Username  string `gorm:"type:varchar(50);not null;unique"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Phone     uint   `gorm:"not null"`
}

type Group struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(50);not null"`
	Adminid   uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}

type Groupmember struct {
	gorm.Model
	Groupid uint `gorm:"not null"`
	Userid  uint `gorm:"not null"`
	Name    string
}

type Expense struct {
	gorm.Model
	Groupid    uint `gorm:"not null"`
	Splitowner uint
	Title      string  `gorm:"type:varchar(255);not null"`
	Place      string  `gorm:"type:varchar(255);not null"`
	Amount     float64 `gorm:"not null"`
	Status     bool
	CreatedAt  time.Time `gorm:"default:now()"`
}

type Split struct {
	gorm.Model
	Expenseid   uint    `gorm:"not null"`
	Userid      uint    `gorm:"not null"`
	Username    string  `gorm:"type:varchar(255);not null"`
	Amount      float64 `gorm:"not null"`
	Paymentid   uint
	Splitstatus bool
}
type Payment struct {
	gorm.Model
	Expenseid uint `gorm:"not null"`
	Splitid   uint `gorm:"not null"`
	Amount    uint `gorm:"not null"`
}

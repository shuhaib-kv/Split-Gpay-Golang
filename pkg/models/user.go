package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `gorm:"type:varchar(50);not null;unique"`
	Lastname  string `gorm:"type:varchar(50);not null;unique"`
	Username  string `gorm:"type:varchar(50);not null;unique"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	Phone     uint   `gorm:"not null"`
}

type Group struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}

type GroupMember struct {
	gorm.Model
	GroupID uint `gorm:"not null"`
	UserID  uint `gorm:"not null"`
}

type Expense struct {
	gorm.Model
	GroupID   uint      `gorm:"not null"`
	Dis       string    `gorm:"type:varchar(255);not null"`
	Amount    float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}

type Split struct {
	gorm.Model
	ExpenseID uint    `gorm:"not null"`
	UserID    uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
}

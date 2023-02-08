package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
	"gorm.io/gorm"
)

func CreateGroup(c *gin.Context) {

	useremail := c.GetString("user")
	fmt.Println(useremail)
	var UsersID int
	err := db.DBS.Raw("select id from users where email=?", useremail).Scan(&UsersID)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user coudnt find",
		})
	}
	var body struct {
		Name string
	}
	c.Bind(&body)

	var Group = models.Group{
		Name:    body.Name,
		Adminid: uint(UsersID),
	}
	db.DBS.Create(&Group)
}
func AddPeoples(c *gin.Context) {

	useremail := c.GetString("user")
	fmt.Println(useremail)
	var UsersID int
	err := db.DBS.Raw("select id from users where email=?", useremail).Scan(&UsersID)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user coudnt find",
		})
	}
	var groupmember models.Groupmember
	fmt.Print(groupmember.Userid)
	if err := c.BindJSON(&groupmember); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	fmt.Print(groupmember.Userid)

	var user []models.User
	for _, i := range user {
		if i.ID != groupmember.Userid {
			c.JSON(http.StatusConflict, gin.H{
				"status": false,
				"error":  " usernot found",
				"data":   "null",
			})
			return
		}
	}
	var group []models.Group
	for _, i := range group {
		if i.Adminid != uint(UsersID) {
			c.JSON(http.StatusConflict, gin.H{
				"status": false,
				"error":  " you are not admin",
				"data":   "null",
			})
			return
		}
	}
	db.DBS.Create(&groupmember)
}

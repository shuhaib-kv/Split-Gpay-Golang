package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func CreateGroup(c *gin.Context) {
	id := c.GetUint("id")
	var body struct {
		Name string
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var Group = models.Group{
		Name:    body.Name,
		Adminid: id,
	}
	if err := db.DBS.Create(&Group).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": false,
			"error":  err.Error(),
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  true,
		"massage": "created group",
		"data":    Group,
	})
}
func AddPeoples(c *gin.Context) {
	id := c.GetUint("id")
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
		if i.Adminid != id {
			c.JSON(http.StatusConflict, gin.H{
				"status": false,
				"error":  " you are not admin of the group",
				"data":   "null",
			})
			return
		}
	}
	db.DBS.Create(&groupmember)
}

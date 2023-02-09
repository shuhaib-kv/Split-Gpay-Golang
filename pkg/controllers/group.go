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
	var groupmember = models.Groupmember{
		Groupid: Group.ID,
		Userid:  id,
	}
	db.DBS.Create(&groupmember)

}
func AddPeoples(c *gin.Context) {
	id := c.GetUint("id")
	var groupmember models.Groupmember
	if err := c.BindJSON(&groupmember); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var group models.Group
	if err := db.DBS.First(&group, "id=?", groupmember.Groupid); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Group Doesn't exist",
			"error":   "error please enter valid information",
		})
		return
	}
	if err := db.DBS.First(&group, "adminid=?", id); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "You are not admin",
			"error":   "error please enter valid information",
		})
		return
	}

	var user models.User
	if err := db.DBS.First(&user, "id=?", groupmember.Userid); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Cant find user",
			"error":   "error please enter valid information",
		})
		return
	}
	// if err := db.DBS.First(&groups, "id=? and userid=?", groupmember.Groupid, groupmember.Userid); err.Error != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"status":  false,
	// 		"message": "user already in the group",
	// 		"error":   "error please ",
	// 	})
	// 	return
	// }
	db.DBS.Create(&groupmember)
	c.JSON(http.StatusAccepted, gin.H{
		"status":  true,
		"message": "Added to success",
		"data":    user.Firstname,
	})
}
func ViewMygroup(c *gin.Context) {
	id := c.GetUint("id")
	var group models.Group
	var groupmember []models.Groupmember
	if err := db.DBS.Find(&groupmember, "userid=?", id).Scan(&groupmember); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Group Doesn't exist",
			"error":   "error please enter valid information",
		})
		return
	}
	for _, i := range groupmember {
		db.DBS.Find(&group, "id=?", i.ID).Scan(&groupmember)
		c.JSON(http.StatusAccepted, gin.H{
			"status":  true,
			"message": "Your Groups",
			"data":    group,
		})
	}

}
func ViewMygroupMembersbyid(c *gin.Context) {
	id := c.GetUint("id")
	var gid uint
	c.BindJSON(&gid)

	fmt.Println(id)
}

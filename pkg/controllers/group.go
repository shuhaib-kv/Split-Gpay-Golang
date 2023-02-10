package controllers

import (
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
			"message": "Group not found",
			"error":   "error please enter valid information",
		})
		return
	}

	if group.Adminid != id {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": "You are not the admin of this group",
			"error":   "error please enter valid information",
		})
		return
	}

	var user models.User
	if err := db.DBS.First(&user, "id=?", groupmember.Userid); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "User not found",
			"error":   "error please enter valid information",
		})
		return
	}

	var isMember models.Groupmember
	if err := db.DBS.First(&isMember, "groupid=? and userid=?", groupmember.Groupid, groupmember.Userid).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "User already in the group",
			"error":   "error please enter valid information",
		})
		return
	}

	groupmember.Name = user.Username
	db.DBS.Create(&groupmember)
	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "User added to the group successfully",
		"data":    user.Firstname,
	})
}

func ViewMygroup(c *gin.Context) {

	id := c.GetUint("id")

	var groupMembers []models.Groupmember
	db.DBS.Where("userid = ?", id).Find(&groupMembers)

	if len(groupMembers) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "You are not in any group",
			"data":    nil,
		})
	} else {
		var groups []models.Group
		groupIDs := make([]uint, len(groupMembers))

		for i, groupMember := range groupMembers {
			groupIDs[i] = groupMember.Groupid
		}

		db.DBS.Where("id IN (?)", groupIDs).Find(&groups)

		groupData := make([]map[string]interface{}, len(groups))
		for i, group := range groups {
			groupData[i] = map[string]interface{}{
				"groupid": group.ID,
				"name":    group.Name,
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Groups retrieved successfully",
			"data":    groupData,
		})
	}
}

func ViewMembers(c *gin.Context) {
	var body struct {
		Group uint `json:"group"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var groupMembers []models.Groupmember
	if err := db.DBS.Where("groupid = ?", body.Group).Find(&groupMembers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Error fetching group members",
			"error":   err.Error(),
		})
		return
	}

	var members []string
	for _, member := range groupMembers {
		var user models.User
		if err := db.DBS.Where("id = ?", member.Userid).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Error fetching user",
				"error":   err.Error(),
			})
			return
		}

		members = append(members, user.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Group members",
		"data":    members,
	})
}

package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

type userid struct {
	ID     uint
	Amount uint
}
type UserID userid

func (u *UserID) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID     uint `json:"id"`
		Amount uint `json:"Amount"`
	}{
		ID:     u.ID,
		Amount: u.Amount,
	})
}

type body struct {
	GroupID uint
	Title   string
	Place   string
	Amount  float64
	Users   []UserID
}

func CreateSplit(c *gin.Context) {

	var expense body

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var sum uint
	for _, userID := range expense.Users {
		var groupmember models.Groupmember
		if err := db.DBS.Where("groupid = ? AND userid=?", userID.ID, expense.GroupID).First(&groupmember).Error; err != nil {

			c.JSON(400, gin.H{
				"error": "not a group member",
				"data":  userID.ID,
			})
			return

		}

		sum += userID.Amount
	}
	if sum != uint(expense.Amount) {
		c.JSON(400, gin.H{
			"data": "error",
			"sum":  "partition doesnt match",
		})
		return

	}
	c.JSON(200, gin.H{
		"data": expense,
		"sum":  sum,
	})

}

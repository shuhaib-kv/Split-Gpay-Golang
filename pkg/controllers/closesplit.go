package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func CloseSplit(c *gin.Context) {
	id := c.GetUint("id")

	var body struct {
		Splitid uint `json:"splitid"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var split models.Split
	if err := db.DBS.Where("id = ?", body.Splitid).First(&split).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Split not found",
			"data":   nil,
		})
		return
	}

	var expense models.Expense
	if err := db.DBS.Where("id = ?", split.Expenseid).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Expense not found",
			"data":   nil,
		})
		return
	}

	if expense.Splitowner != id {
		c.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"error":  "You are not the split owner",
			"data":   nil,
		})
		return
	}

	split.Splitstatus = true
	if err := db.DBS.Save(&split).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Error while updating the split status",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Split closed",
		"data": gin.H{
			"split_status": split.Splitstatus,
			"user_id":      split.Userid,
		},
	})
}
func CloseExpense(c *gin.Context) {

	var body struct {
		Expenceid uint `json:"expenceid"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	id := c.GetUint("id")
	var expense models.Expense
	if err := db.DBS.Find(&expense, body.Expenceid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "Expense not found",
			"data":   "null",
		})
		return
	}
	if expense.Splitowner != id {
		c.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"error":  "Only the split owner can close the expense",
			"data":   "null",
		})
		return
	}
	if err := db.DBS.Model(&expense).Update("status", true).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Failed to update expense status",
			"data":   "null",
		})
		return
	}
	var splits []models.Split
	if err := db.DBS.Where("expenseid = ?", body.Expenceid).Find(&splits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  "Failed to retrieve splits for the expense",
			"data":   "null",
		})
		return
	}
	for _, split := range splits {
		if split.Splitstatus == false {
			if err := db.DBS.Model(&split).Update("splitstatus", true).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": false,
					"error":  "Failed to update split status",
					"data":   "null",
				})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Expense and its splits closed successfully",
		"data":    "null",
	})

}

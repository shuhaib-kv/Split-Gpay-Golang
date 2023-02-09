package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func ViewSplit(c *gin.Context) {
	var body struct {
		expenceid uint
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var expense models.Expense
	var split []models.Split
	if err := db.DBS.First(&expense, "id=1"); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Group Doesn't exist",
			"error":   "error please enter valid information",
		})
		return
	}
	db.DBS.Find(&split, "expenseid=?", expense.ID).Scan(&split)
	for _, i := range split {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Your Groups",
			"data": gin.H{
				"Group id":       i.ID,
				"Group name":     i.Username,
				"amount ":        i.Amount,
				"payment status": i.Paymentstatus,
				"split status":   i.Splitstatus,
			},
		})
	}

	// c.JSON(http.StatusFound, gin.H{
	// 	"data": expense,
	// })
	return
}
func ViewMysplit(c *gin.Context) {
	id := c.GetUint("id")
	var gid struct {
		groupid uint
	}
	if err := c.BindJSON(&gid); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var expense []models.Expense
	db.DBS.Find(&expense, "groupid=? and spliowner=? ", gid.groupid, id).Scan(&expense)
	for _, i := range expense {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Your Groups",
			"data": gin.H{
				"expenceid": i.ID,
				"tittle":    i.Title,
				"amount ":   i.Amount,
			},
		})
	}
}

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
	var splitid models.Split
	db.DBS.Find(&splitid).Where(splitid.ID == body.Splitid).Scan(&splitid)
	var expense models.Expense
	db.DBS.Find(&expense).Where(expense.ID == splitid.Expenseid).Scan(&expense)
	if expense.Splitowner != id {
		c.JSON(http.StatusConflict, gin.H{
			"status":  true,
			"massage": "you are not admin",
			"data":    "onliy split owner can close expense",
		})
		return
	}

}
func CloseExpense(c *gin.Context) {
	id := c.GetUint("id")

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
	var expense models.Expense
	db.DBS.Find(&expense).Where(expense.ID == body.Expenceid).Scan(&expense)
	if expense.Splitowner != id {
		c.JSON(http.StatusConflict, gin.H{
			"status":  true,
			"massage": "you are not admin",
			"data":    "onliy split owner can close expense",
		})
		return
	}

}

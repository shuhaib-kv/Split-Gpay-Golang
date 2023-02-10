package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func PaySplit(c *gin.Context) {
	var body struct {
		amount    uint
		expenceid uint
		splitid   uint
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	var expence models.Expense
	var Split models.Split

	db.DBS.Find(&expence, "id=?", body.expenceid).Scan(&expence)
	db.DBS.Find(&Split, "id=? expenseid", body.expenceid).Scan(&Split)
	if expence.Status == true {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "split closed",
			"data":   expence.Status,
		})
		return
	}
	if Split.Splitstatus == true {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Your split closed by admin",
			"data":   expence.Status,
		})
		return
	}
	if Split.Paymentstatus == true {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Your paid the split",
			"data":   Split.Paymentstatus,
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"status": false,
		"error":  "Your paid the split",
		"data":   Split.Paymentstatus})

}

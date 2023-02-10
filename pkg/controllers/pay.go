package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func PaySplit(c *gin.Context) {
	var body struct {
		Amount    uint `json:"amount"`
		Expenceid uint `json:"expenceid"`
		Splitid   uint `json:"splitid"`
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

	db.DBS.Find(&expence, "id=?", body.Expenceid).Scan(&expence)
	db.DBS.Find(&Split, "expenseid=? ", body.Expenceid).Scan(&Split)
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
	if Split.Amount != float64(body.Amount) {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Amount doesnt match",
			"data":   Split.Paymentstatus,
		})
		return
	}
	var pay = models.Payment{
		Expenseid: body.Expenceid,
		Splitid:   body.Splitid,
		Amount:    body.Amount,
	}
	db.DBS.Create(&pay)
	c.JSON(http.StatusFound, gin.H{
		"status": true,
		"error":  "Your paid the split",
		"data":   pay})
	var done = models.Split{
		Paymentstatus: true,
		Paymentid:     pay.ID,
		Splitstatus:   true,
	}
	db.DBS.Model(&Split).Where("splits.id=?", pay.Splitid).Updates(&done)
	db.DBS.Raw("select *  from split where expenseid=? and paymentstatus=?", body.Expenceid, true)
	var splitAmount float64

	db.DBS.Table("splits").Where("expenseid = ?", body.Expenceid).Select("SUM(amount)").Row().Scan(&splitAmount)

	difference := expence.Amount - splitAmount
	if difference != 0 {

	}

}

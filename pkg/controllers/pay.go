package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func PaySplit(c *gin.Context) {
	id := c.GetUint("id")

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
	db.DBS.Find(&Split, "expenseid=? and id=? and userid=?", body.Expenceid, body.Splitid, id).Scan(&Split)
	if expence.Status == true {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "split closed",
			"data":   expence.Status,
		})
		return
	}
	if Split.Splitstatus == true && expence.Status == true {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Your split closed by admin",
			"data":   expence.Status,
		})
		return
	}
	if Split.Splitstatus == true {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Your paid the split",
			"data":   "",
		})
		return
	}
	if Split.Amount != float64(body.Amount) {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Amount doesnt match",
			"data":   Split.Amount,
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
	// var done = models.Split{

	// 	Splitstatus: true,
	// }
	db.DBS.Model(&Split).Update("splitstatus=?", true).Where("splits.id=?", body.Splitid)
	db.DBS.Model(&Split).Update("paymentid=?", pay.ID).Where("splits.id=?", body.Splitid)

}

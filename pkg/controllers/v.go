package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func ViewExpense(c *gin.Context) {
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
	var expence []models.Expense
	db.DBS.Find(&expence, "groupid=?", body.Expenceid).Order("expenses.created_at DESC").Scan(&expence)
	expencedata := make([]map[string]interface{}, len(expence))
	for i, s := range expence {
		expencedata[i] = map[string]interface{}{
			"id":     s.ID,
			"title":  s.Title,
			"place":  s.Place,
			"amount": s.Amount,
			"status": s.Status,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": " group expences",
		"data":    expencedata,
	})

}
func ViewExpenseNotClosed(c *gin.Context) {
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
	var expence []models.Expense
	db.DBS.Find(&expence, "groupid=? and status=?", body.Expenceid, false).Order("expenses.created_at DESC").Scan(&expence)
	expencedata := make([]map[string]interface{}, len(expence))
	for i, s := range expence {
		expencedata[i] = map[string]interface{}{
			"id":     s.ID,
			"title":  s.Title,
			"place":  s.Place,
			"amount": s.Amount,
			"status": s.Status,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": " group expences",
		"data":    expencedata,
	})

}
func ViewExpenseClosed(c *gin.Context) {
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
	var expence []models.Expense
	db.DBS.Find(&expence, "groupid=? and status=?", body.Expenceid, true).Order("expenses.created_at DESC").Scan(&expence)
	expencedata := make([]map[string]interface{}, len(expence))
	for i, s := range expence {
		expencedata[i] = map[string]interface{}{
			"id":     s.ID,
			"title":  s.Title,
			"place":  s.Place,
			"amount": s.Amount,
			"status": s.Status,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": " group expences",
		"data":    expencedata,
	})

}

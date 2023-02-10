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
	splitData := make([]map[string]interface{}, len(split))
	for i, s := range split {
		splitData[i] = map[string]interface{}{
			"split id":     s.ID,
			"user name":    s.Username,
			"amount":       s.Amount,
			"split status": s.Splitstatus,
		}
	}

	c.JSON(200, gin.H{
		"status":  true,
		"message": "your split",
		"data":    splitData,
	})
	return
}
func ViewMysplit(c *gin.Context) {
	userID := c.GetUint("id")

	// Get the group ID from the JSON request body
	var groupID struct {
		GroupID int
	}
	if err := c.BindJSON(&groupID); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}

	// Retrieve the expenses from the database that match the user ID and group ID
	var expenses []models.Expense
	db.DBS.Find(&expenses, "groupid = ? and splitowner = ?", groupID.GroupID, userID).Scan(&expenses)

	// If there are no matching expenses, return a 404 status code
	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "No matching expenses found",
			"data":   "null",
		})
		return
	}

	// Build the response data
	var responseData []gin.H
	for _, expense := range expenses {
		// Retrieve the splits for each expense
		var splits []models.Split
		db.DBS.Find(&splits, "expenseid = ?", expense.ID).Scan(&splits)

		// Build the expense data
		expenseData := gin.H{
			"expenseID": expense.ID,
			"title":     expense.Title,
			"amount":    expense.Amount,
			"splits":    splits,
			"status":    expense.Status,
			"createdAt": expense.CreatedAt,
		}
		responseData = append(responseData, expenseData)
	}

	// Return the response data
	c.JSON(200, gin.H{
		"status":  true,
		"message": "Split details retrieved successfully",
		"data":    responseData,
	})

}
func ViewWhoNotPaid(c *gin.Context) {
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
	db.DBS.Find(&split, "expenseid=? and splitstatus=?", expense.ID, false).Scan(&split)
	splitData := make([]map[string]interface{}, len(split))
	for i, s := range split {
		splitData[i] = map[string]interface{}{
			"split id":     s.ID,
			"userid":       s.Userid,
			"split owner":  s.Username,
			"amount":       s.Amount,
			"split status": s.Splitstatus,
		}
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "persons who not paid",
		"data":    splitData,
	})
}

func ViewWhoPaid(c *gin.Context) {
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
	db.DBS.Find(&split, "expenseid=? and splitstatus=?", expense.ID, true).Scan(&split)

	splitData := make([]map[string]interface{}, len(split))
	for i, s := range split {
		splitData[i] = map[string]interface{}{
			"split id":     s.ID,
			"userid":       s.Userid,
			"split owner":  s.Username,
			"amount":       s.Amount,
			"split status": s.Splitstatus,
		}
	}
	c.JSON(200, gin.H{
		"status":  true,
		"message": "persons who  paid",
		"data":    splitData,
	})
}

package controllers

import (
	"encoding/json"
	"net/http"

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
	id := c.GetUint("id")
	var expense body
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var sum uint
	for _, userID := range expense.Users {
		var groupmember models.Groupmember
		if err := db.DBS.Where("groupid = ? AND userid=?", expense.GroupID, userID.ID).First(&groupmember).Error; err != nil {

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
	var expensedb = models.Expense{
		Groupid:    expense.GroupID,
		Splitowner: id,
		Title:      expense.Title,
		Place:      expense.Place,
		Amount:     expense.Amount,
	}
	db.DBS.Create(&expensedb)
	var user models.User

	for _, i := range expense.Users {
		db.DBS.Where("id = ? ", i.ID).Find(&user)
		var split = models.Split{
			Userid:      i.ID,
			Username:    user.Username,
			Amount:      float64(i.Amount),
			Expenseid:   expensedb.ID,
			Splitstatus: false,
		}
		db.DBS.Create(&split)

	}

	c.JSON(200, gin.H{
		"data": expense,
		"sum":  sum,
	})

}
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

	var expenses []models.Expense
	db.DBS.Find(&expenses, "groupid = ? and splitowner = ?", groupID.GroupID, userID).Scan(&expenses)

	if len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  "No matching expenses found",
			"data":   "null",
		})
		return
	}

	var responseData []gin.H
	for _, expense := range expenses {
		var splits []models.Split
		db.DBS.Find(&splits, "expenseid = ?", expense.ID).Scan(&splits)

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

	c.JSON(200, gin.H{
		"status":  true,
		"message": "Split details retrieved successfully",
		"data":    responseData,
	})

}

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

package controllers

// func ViewSplitExpencesInGroup(c *gin.Context) {
// 	// id := c.GetUint("id")
// 	var body struct {
// 		groupid uint
// 	}
// 	c.BindJSON(&body)
// 	var expense []models.Expense

// 	if err := db.DBS.Find(&expense, "groupid=?", body.groupid).Scan(&expense); err.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "Group Doesn't exist",
// 			"error":   "error please enter valid information",
// 		})
// 		return
// 	}
// 	for _, i := range expense {

// 		c.JSON(200, gin.H{
// 			"wow": i.Splitowner,
// 		})
// 	}
// 	c.JSON(200, gin.H{
// 		"wow": expense,
// 	})

// }
// func CloseAllSplit(c *gin.Context) {

// }
// func CloseSplitByPerson(c *gin.Context) {}
// func PaySplit(c *gin.Context)           {}
// func ViewsplitDetails(c *gin.Context) {
// 	var body struct {
// 		expid   uint
// 		groupid uint
// 	}
// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(http.StatusConflict, gin.H{
// 			"status": false,
// 			"error":  "Invalid JSON",
// 			"data":   "null",
// 		})
// 		return
// 	}

// 	var expense models.Expense
// 	var split models.Split
// 	if err := db.DBS.Find(&expense, "id=?", body.expid).Scan(&expense); err.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "Group Doesn't exist",
// 			"error":   "error please enter valid information",
// 		})
// 		return
// 	}
// 	if err := db.DBS.Find(&split, "expenseid=?", body.expid).Scan(&split); err.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"status":  false,
// 			"message": "Group Doesn't exist",
// 			"error":   "error please enter valid information",
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{

// 		"data": split,
// 	})
// }

// func ViewNotPaid(c *gin.Context)         {}
// func ViewPaid(c *gin.Context)            {}
// func ViewAllsplitInGroup(c *gin.Context) {}

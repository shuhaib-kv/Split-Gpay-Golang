package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
	if err := db.DBS.Create(&user).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": false,
			"error":  err.Error(),
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":  true,
		"massage": "created user",
		"data":    user,
	})
}

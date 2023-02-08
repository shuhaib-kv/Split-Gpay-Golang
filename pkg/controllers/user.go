package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/models"
)

func SignUp(c gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}
	//		if err := database.db.Create(&user).Error; err != nil {
	//			c.JSON(500, gin.H{"error": err.Error()})
	//			return
	//		}
	//		c.JSON(200, gin.H{"user": user})
	//	}
}

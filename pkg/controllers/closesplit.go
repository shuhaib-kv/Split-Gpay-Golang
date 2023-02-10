package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CloseBuyID(c *gin.Context) {
	var body struct {
		APlitid uint `json:"expenceid"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": false,
			"error":  "Invalid JSON",
			"data":   "null",
		})
		return
	}
}

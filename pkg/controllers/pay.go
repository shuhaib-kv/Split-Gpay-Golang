package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusFound, gin.H{
		"": "",
	})

}

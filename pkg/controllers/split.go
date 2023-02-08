package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateSplit(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": s,
	})

}

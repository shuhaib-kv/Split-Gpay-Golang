package controllers

import "github.com/gin-gonic/gin"

func CloseBuyID(c *gin.Context) {
	var body struct {
		splitid int
		groupid int
	}
	c.BindJSON(&body)
}

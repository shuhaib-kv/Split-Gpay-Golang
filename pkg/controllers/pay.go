package controllers

import "github.com/gin-gonic/gin"

func PaySplit(c *gin.Context) {
	var body struct {
		amount    uint
		expenceid uint
		splitid   uint
	}
	c.BindJSON(&body)

}

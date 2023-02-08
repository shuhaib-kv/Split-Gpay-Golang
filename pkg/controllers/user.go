package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/middleware"
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
func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "binding json faild",
			"error":   "error ",
		})
		return
	}
	if body.Email == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status":  false,
			"message": "Email is required",
			"error":   "error",
		})
		return
	}
	if body.Password == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status":  false,
			"message": "Password is required",
			"error":   "error",
		})

		return
	}

	var user models.User
	if err := db.DBS.First(&user, "email = ?", body.Email); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Cant find user",
			"error":   "error please enter valid information",
		})
		return
	}
	if user.Password != body.Password {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status":  false,
			"message": "Incorrect Password",
			"error":   "error please enter valid information",
		})
		return
	}
	if user.Email != body.Email {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status":  false,
			"message": "Incorrect email",
			"error":   "error please enter valid information",
		})
		return
	}
	tokenString, err := middleware.GenerateJWT(user.Email, user.ID)
	c.SetCookie("UserAuth", tokenString, 3600*24*30, "", "", false, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "ok",
		"data":    tokenString,
	})

}
func Home(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "ok",
		"data":    "",
	})
}

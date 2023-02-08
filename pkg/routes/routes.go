package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/controllers"
)

func RoutesOfApi(c *gin.Engine) {
	c.POST("/user", controllers.SignUp)
}

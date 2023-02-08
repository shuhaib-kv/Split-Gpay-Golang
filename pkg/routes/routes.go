package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/controllers"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/middleware"
)

func RoutesOfApi(c *gin.Engine) {
	c.POST("/user/signin", controllers.SignUp)
	c.GET("/user/login", controllers.Login)
	c.GET("/user/home", middleware.UserAuth(), controllers.Home)
}

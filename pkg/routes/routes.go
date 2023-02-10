package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/controllers"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/middleware"
)

func RoutesOfApi(c *gin.Engine) {

	c.POST("/user/signup", controllers.SignUp)
	c.GET("/user/login", controllers.Login)
	c.GET("/user/home", middleware.UserAuth, controllers.Home)
	c.POST("/user/group/create", middleware.UserAuth, controllers.CreateGroup)
	c.POST("/user/group/add", middleware.UserAuth, controllers.AddPeoples)
	c.GET("/user/group/view", middleware.UserAuth, controllers.ViewMygroup)
	c.GET("/user/group/view/mygroups", middleware.UserAuth, controllers.ViewMygroupMembersbyid)
	c.POST("/user/group/split", middleware.UserAuth, controllers.CreateSplit)
	c.GET("/user/view/split", middleware.UserAuth, controllers.ViewSplit)
	c.GET("/user/view/notpaid", middleware.UserAuth, controllers.ViewWhoNotPaid)
	c.GET("/user/view/paid", middleware.UserAuth, controllers.ViewWhoPaid)
	c.POST("/user/pay/split", middleware.UserAuth, controllers.PaySplit)
}

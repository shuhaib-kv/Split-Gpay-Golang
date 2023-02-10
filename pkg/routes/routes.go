package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/controllers"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/middleware"
)

func RoutesOfApi(c *gin.Engine) {

	c.POST("/user/signup", controllers.SignUp)                                 //done
	c.GET("/user/login", controllers.Login)                                    //done
	c.GET("/user/home", middleware.UserAuth, controllers.Home)                 //done
	c.POST("/user/group/create", middleware.UserAuth, controllers.CreateGroup) //done
	c.POST("/user/group/add", middleware.UserAuth, controllers.AddPeoples)     //done
	c.GET("/user/group/view", middleware.UserAuth, controllers.ViewMygroup)    //done
	c.POST("/user/group/split", middleware.UserAuth, controllers.CreateSplit)  //done

	c.GET("/user/view/split", middleware.UserAuth, controllers.ViewSplit)
	c.GET("/user/view/mysplit", middleware.UserAuth, controllers.ViewMysplit)
	c.GET("/user/group/view/members", middleware.UserAuth, controllers.ViewMembers)
	c.GET("/user/view/notpaid", middleware.UserAuth, controllers.ViewWhoNotPaid)
	c.GET("/user/view/paid", middleware.UserAuth, controllers.ViewWhoPaid)

	c.POST("/user/pay/split", middleware.UserAuth, controllers.PaySplit)
	//close split
	c.PATCH("/user/group/split/close/individual")
}

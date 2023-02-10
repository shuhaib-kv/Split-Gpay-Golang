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

	c.GET("/user/view/expenses", middleware.UserAuth, controllers.ViewSplit)        //done
	c.GET("/user/view/mysplit", middleware.UserAuth, controllers.ViewMysplit)       //done
	c.GET("/user/group/view/members", middleware.UserAuth, controllers.ViewMembers) //done
	c.GET("/user/view/notpaid", middleware.UserAuth, controllers.ViewWhoNotPaid)    //done
	c.GET("/user/view/paid", middleware.UserAuth, controllers.ViewWhoPaid)          //done

	c.POST("/user/pay/split", middleware.UserAuth, controllers.PaySplit) //done

	c.PATCH("/user/group/split/close", middleware.UserAuth, controllers.CloseSplit)
	c.PATCH("/user/group/expense/close", middleware.UserAuth, controllers.CloseExpense)

	c.GET("/user/view/expense", middleware.UserAuth, controllers.ViewExpense)                     //done
	c.GET("/user/view/expense/closed/not", middleware.UserAuth, controllers.ViewExpenseNotClosed) //done
	c.GET("/user/view/expense/closed", middleware.UserAuth, controllers.ViewExpenseClosed)        //done

}

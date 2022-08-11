package router

import "github.com/gin-gonic/gin"

func (c *CapsuleRouter) UserRouter(r *gin.RouterGroup) {
	userController := c.Controller.UserController

	r.POST("/user/login", userController.LoginUser)

	r.POST("/user/signup", userController.CreateUser)

	r.POST("/user/logout", userController.LogoutUser)
}

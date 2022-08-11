package router

import "github.com/gin-gonic/gin"

func (c *CapsuleRouter) RoleRouter(r *gin.RouterGroup) {
	roleController := c.Controller.RoleController

	r.GET("/role", roleController.GetAllRoles)

	r.GET("/role/:id", roleController.GetRoleById)

	r.POST("/role", roleController.CreateRole)
}

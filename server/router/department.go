package router

import "github.com/gin-gonic/gin"

func (c *CapsuleRouter) DepartmentRouter(r *gin.RouterGroup) {
	departmentController := c.Controller.DepartmentController

	r.GET("/department", departmentController.GetAllDepartments)

	r.GET("/department/:id", departmentController.GetDepartmentById)

	r.POST("/department", departmentController.CreateDepartment)

	r.PUT("/department/:id", departmentController.UpdateDepartment)
}

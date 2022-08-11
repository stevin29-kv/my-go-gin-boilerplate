package router

import "github.com/gin-gonic/gin"

func (c *CapsuleRouter) EmployeeRouter(r *gin.RouterGroup) {
	employeeController := c.Controller.EmployeeController

	r.GET("/employee", employeeController.GetAllEmployees)

	r.GET("/employee/:id", employeeController.GetEmployeeById)

	r.POST("/employee", employeeController.CreateEmployee)

	r.DELETE("/employee/:id", employeeController.DeleteEmployee)

	r.PUT("/employee/:id", employeeController.UpdateEmployee)

	r.PATCH("/employee/:id", employeeController.UpdateEmployeeStatusById)

	r.PATCH("/employee/id-proof/:id", employeeController.UploadIdProof)
}

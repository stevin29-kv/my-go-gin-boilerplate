package controller

import "employee-app/internal/service"

type Controller struct {
	HealthController     *healthController
	EmployeeController   *employeeController
	DepartmentController *departmentController
	RoleController       *roleController
	UserController       *userController
}

func InitController(s *service.Service) *Controller {
	return &Controller{
		HealthController:     InitHealthController(),
		EmployeeController:   InitEmployeeController(s.Employee),
		DepartmentController: InitDepartmentController(s.Department),
		RoleController:       InitRoleController(s.Role),
		UserController:       InitUserController(s.User),
	}
}

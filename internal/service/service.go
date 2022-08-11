package service

import (
	"employee-app/internal/repository"
	"employee-app/internal/service/department"
	"employee-app/internal/service/employee"
	"employee-app/internal/service/role"
	"employee-app/internal/service/user"
)

type Service struct {
	Employee   employee.ServiceInterface
	Department department.ServiceInterface
	Role       role.ServiceInterface
	User       user.ServiceInterface
}

func InitService(repo *repository.Repository) *Service {
	return &Service{
		Employee: employee.InitEmployeeService(
			repo.Employee,
		),
		Department: department.InitDepartmentService(
			repo.Department,
		),
		Role: role.InitRoleService(
			repo.Role,
		),
		User: user.InitUserService(
			repo.User,
		),
	}
}

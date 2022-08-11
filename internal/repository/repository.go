package repository

import (
	"employee-app/internal/repository/department"
	"employee-app/internal/repository/employee"
	"employee-app/internal/repository/role"
	"employee-app/internal/repository/user"

	"gorm.io/gorm"
)

type Repository struct {
	Employee   employee.RepositoryInterface
	Department department.RepositoryInterface
	Role       role.RepositoryInterface
	User       user.RepositoryInterface
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		Employee:   employee.InitEmployeeRepository(db),
		Department: department.InitDepartmentRepository(db),
		Role:       role.InitiRoleRepository(db),
		User:       user.InitUserRepository(db),
	}
}

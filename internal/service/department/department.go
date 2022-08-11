package department

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/model"
	"employee-app/internal/repository/department"

	"gorm.io/gorm"
)

type ServiceInterface interface {
	CreateDepartment(createDepartmentDto dto.CreateDepartment) *model.APIResponse
	GetAllDepartments() *model.APIResponse
	GetDepartmentById(id string) *model.APIResponse
	UpdateDepartment(UpdateDepartmentDto dto.UpdateDepartment, id string) *model.APIResponse
}

type service struct {
	department department.RepositoryInterface
	db         *gorm.DB
}

func InitDepartmentService(department department.RepositoryInterface, db *gorm.DB) ServiceInterface {
	return &service{
		department: department,
		db:         db,
	}
}

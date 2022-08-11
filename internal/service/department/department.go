package department

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/model"
	"employee-app/internal/repository/department"
)

type ServiceInterface interface {
	CreateDepartment(createDepartmentDto dto.CreateDepartment) *model.APIResponse
	GetAllDepartments() *model.APIResponse
	GetDepartmentById(id string) *model.APIResponse
	UpdateDepartment(UpdateDepartmentDto dto.UpdateDepartment, id string) *model.APIResponse
}

type service struct {
	department department.RepositoryInterface
}

func InitDepartmentService(department department.RepositoryInterface) ServiceInterface {
	return &service{
		department: department,
	}
}

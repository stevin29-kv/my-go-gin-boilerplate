package employee

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/model"
	"employee-app/internal/repository/employee"
)

type ServiceInterface interface {
	CreateEmployee(employeeRequest dto.CreateEmployeeRequest) *model.APIResponse
	GetAllEmployees(string, string, string) *model.APIResponse
	GetEmployeeById(string) *model.APIResponse
	DeleteEmployee(string) *model.APIResponse
	UpdateEmployee(string, dto.UpdateEmployeeRequest) *model.APIResponse
	UpdateEmployeeStatusById(string, dto.UpdateEmployeeStatusRequest) *model.APIResponse
	UploadIdProof(id string, newFileName string) *model.APIResponse
}

type service struct {
	employee employee.RepositoryInterface
}

func InitEmployeeService(employee employee.RepositoryInterface) ServiceInterface {
	return &service{
		employee: employee,
	}
}

package employee

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/helpers"
	"employee-app/internal/entity/model"
	"employee-app/logger"
	"strconv"
)

func (es *service) CreateEmployee(employeeRequest dto.CreateEmployeeRequest) *model.APIResponse {
	logger.Info("Start CreateEmployee in Service")
	address := model.Address{
		Street: employeeRequest.Street,
		City:   employeeRequest.City,
		State:  employeeRequest.State,
	}

	employee := model.Employee{
		Name:         employeeRequest.Name,
		Username:     employeeRequest.Username,
		Email:        employeeRequest.Email,
		Age:          employeeRequest.Age,
		IsActive:     true,
		DepartmentID: employeeRequest.DepartmentID,
		RoleID:       employeeRequest.RoleID,
		Address:      address,
	}
	logger.Info(employee)

	employee, err := es.employee.CreateEmployee(employee)
	if err != nil {
		logger.Error("Error while saving product")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot save employee",
			},
		}
	}
	logger.Info("Saved employee")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       employee,
	}
}

func (es *service) GetAllEmployees(search string, sort_by string, order string) *model.APIResponse {
	logger.Info("Start GetAllEmployees in Service")
	filter := helpers.Pagination{
		Filter: search,
		SortBy: sort_by,
		Order:  order,
	}
	employee, err := es.employee.GetAllEmployees(filter)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employees not found",
			},
		}
	}

	logger.Info("End GetAllEmployees in Service")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}

func (es *service) GetEmployeeById(id string) *model.APIResponse {
	logger.Info("Started GetEmployeeById in Service")
	employee, err := es.employee.GetEmployeeById(id)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}

	logger.Info("End GetEmployeeById in Service")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}

func (es *service) DeleteEmployee(id string) *model.APIResponse {
	logger.Info("Start DeleteEmployee in Service")
	employee, err := es.employee.GetEmployeeById(id)
	logger.Info(employee)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}

	err = es.employee.DeleteEmployee(id)
	if err != nil {
		logger.Error("Error while deleting employee")
		return &model.APIResponse{
			StatusCode: 404,
			Data:       "Failed to delete",
		}
	}

	logger.Info("Deleted Employee")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       "Successfully Deleted",
	}
}

func (es *service) UpdateEmployee(id string, employeeRequest dto.UpdateEmployeeRequest) *model.APIResponse {
	logger.Info("Start UpdateEmployee - Service")
	employee, err := es.employee.GetEmployeeById(id)
	logger.Info(employee)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}
	address := model.Address{
		Street: employeeRequest.Street,
		City:   employeeRequest.City,
		State:  employeeRequest.State,
	}
	// logger.Info(address)
	address, err = es.employee.UpdateAddress(strconv.Itoa(employee.AddressID), address)
	if err != nil {
		logger.Error("Error while updating address")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update address",
			},
		}
	}

	employee = model.Employee{
		Name:         employeeRequest.Name,
		Username:     employeeRequest.Username,
		Email:        employeeRequest.Email,
		Age:          employeeRequest.Age,
		IsActive:     employeeRequest.IsActive,
		DepartmentID: employeeRequest.DepartmentID,
		RoleID:       employeeRequest.RoleID,
		Address:      address,
	}
	logger.Info(employee)

	employee, ok := es.employee.UpdateEmployee(id, employee)
	if ok != nil {
		logger.Error("Error while updating employee")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update employee",
			},
		}
	}
	employee, err = es.employee.UpdateEmployeeStatusById(id, employeeRequest.IsActive)
	if err != nil {
		logger.Error("Error while updating status")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update employee",
			},
		}
	}
	logger.Info("Updated employee")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}

func (es *service) UpdateEmployeeStatusById(id string, employeeRequest dto.UpdateEmployeeStatusRequest) *model.APIResponse {
	logger.Info("Start UpdateEmployee - Service")
	employee, err := es.employee.GetEmployeeById(id)
	logger.Info(employee)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Employee not found",
			},
		}
	}

	employee, err = es.employee.UpdateEmployeeStatusById(id, employeeRequest.IsActive)
	if err != nil {
		logger.Error("Error in service")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Update Failed",
			},
		}
	}
	logger.Info(employee)
	return &model.APIResponse{
		StatusCode: 200,
		Data:       employee,
	}
}

func (es *service) UploadIdProof(id string, newFileName string) *model.APIResponse {
	logger.Info("Start UploadIdProof in Service")
	employeeData, err := es.employee.UploadIdProof(id, newFileName)
	if err != nil {
		logger.Error("Error while updating file name field")
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot update file name",
			},
		}
	}
	logger.Info("End UploadIdProof in Service")
	return &model.APIResponse{
		StatusCode: 201,
		Data:       employeeData,
	}
}

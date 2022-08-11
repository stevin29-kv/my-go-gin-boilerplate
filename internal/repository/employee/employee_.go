package employee

import (
	"employee-app/internal/entity/helpers"
	"employee-app/internal/entity/model"
	"employee-app/logger"

	"gorm.io/gorm"
)

func (er *repository) CreateEmployee(employee model.Employee) (model.Employee, error) {
	logger.Info("Start CreateEmployee in Repo")
	err := er.db.Create(&employee).Preload("Department").Preload("Role").Preload("Address").Preload("Department.Department").First(&employee).Error
	logger.Info("End CreateEmployee in Repo")
	return employee, err
}

func (er *repository) GetAllEmployees(filter helpers.Pagination) ([]model.Employee, error) {
	logger.Info("Start GetAllEmployees in Repo")
	var employee []model.Employee
	if filter.Filter == "" && filter.SortBy == "" {
		err := er.db.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").Find(&employee).Error
		return employee, err
	} else if filter.Filter == "" {
		err := er.db.Order(filter.SortBy + " " + filter.Order).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").Find(&employee).Error
		return employee, err
	} else if filter.SortBy == "" {
		err := er.db.Where("username LIKE ?", filter.Filter+"%").Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").Find(&employee).Error
		return employee, err
	} else {
		err := er.db.Order(filter.SortBy+" "+filter.Order).Where("username LIKE ?", filter.Filter+"%").Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").Find(&employee).Error
		logger.Info("End GetAllEmployees in Repo")
		return employee, err
	}
}

func (er *repository) GetEmployeeById(id string) (model.Employee, error) {
	logger.Info("Started GetEmployeeById in Repo")
	var employee model.Employee
	err := er.db.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "id = ?", id).Error
	logger.Info("Ended GetEmployeeById in Repo")
	return employee, err
}

func (er *repository) CreateAddress(address model.Address) (model.Address, error) {
	logger.Info("Start CreateAddress in Repo")
	err := er.db.Create(&address).Error
	logger.Info("End CreateAddress in Repo")
	return address, err
}

func (er *repository) GetAddressById(id int) (model.Address, error) {
	logger.Info("Started GetAddressById in Repo")
	var address model.Address
	err := er.db.First(&address, "id = ?", id).Error
	logger.Info("Ended GetAddressById in Repo")
	return address, err
}

func (er *repository) DeleteEmployee(id string) error {
	logger.Info("Start DeleteEmployee in Repo")
	var employee model.Employee
	err := er.db.Delete(&employee, "id = ?", id).Error
	logger.Info("End DeleteEmployee in Repo")
	return err
}

func (er *repository) UpdateEmployee(id string, employee model.Employee) (model.Employee, error) {
	logger.Info("Started UpdateEmployee in Repo")
	var employeeData model.Employee
	err := er.db.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employeeData, "id = ?", id).Error
	if err != nil {
		logger.Error("Employee not found")
		return employeeData, err
	}
	logger.Info(employeeData, employee)
	err = er.db.Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", id).Updates(&employee).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "id = ?", id).Error
	logger.Info("Ended UpdateEmployee in Repo")
	return employee, err
}

func (er *repository) UpdateAddress(id string, address model.Address) (model.Address, error) {
	logger.Info("Started UpdateAddress in Repo")
	err := er.db.Where("id = ?", id).Updates(&address).Error
	logger.Info("Ended UpdateAddress in Repo")
	return address, err
}

func (er *repository) UpdateEmployeeStatusById(id string, is_active bool) (model.Employee, error) {
	logger.Info("Started UpdateEmployeeStatusById in Repo")
	var employee model.Employee
	err := er.db.Model(&employee).Where("id = ?", id).Update("is_active", is_active).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "id = ?", id).Error
	logger.Info("Ended UpdateEmployeeStatusById in Repo")
	return employee, err
}

func (er *repository) GetEmployeeByEmail(email string) (model.Employee, error) {
	logger.Info("Started GetEmployeeById in Repo")
	var employee model.Employee
	err := er.db.Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employee, "email = ?", email).Error
	logger.Info("Ended GetEmployeeById in Repo")
	return employee, err
}

func (er *repository) UploadIdProof(id string, newFileName string) (model.Employee, error) {
	logger.Info("Start UploadIdProof in Repo")
	var employeeData model.Employee
	if err := er.db.Where("id = ?", id).First(&employeeData).Error; err != nil {
		logger.Error("Error while fetching from employee repo", err.Error())
		return model.Employee{}, err
	}
	err := er.db.Model(&employeeData).Update("IdProof", newFileName).Preload("Address").Preload("Role").Preload("Department").Preload("Department.Department").First(&employeeData).Error
	logger.Info("End UploadIdProof in Repo")
	return employeeData, err
}

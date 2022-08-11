package department

import (
	"employee-app/internal/entity/model"
	"employee-app/logger"
)

func (r *repository) CreateDepartment(department model.Department) (model.Department, error) {
	logger.Infof("Start CreateDepartment %+v ", department)
	err := r.db.Create(&department).Error
	logger.Info("End CreateDepartment")
	return department, err
}

func (r *repository) GetAllDepartments() ([]model.Department, error) {
	var departments []model.Department
	res := r.db.Preload("Department").Find(&departments)
	if res.Error != nil {
		msg := res.Error
		return nil, msg
	}
	return departments, res.Error
}

func (r *repository) GetDepartmentById(id string) (model.Department, error) {
	logger.Info("Start GetDepartmentById")
	var department model.Department
	response := r.db.Where("id =?", id).Preload("Department").First(&department)
	logger.Info(response.Error)
	if response.Error != nil {
		logger.Error("Error while fetching from department repo", response.Error.Error())
	}
	logger.Infof("End GetDepartmentById")
	return department, response.Error
}

func (r *repository) UpdateDepartment(updatedDepartment model.Department, id string) (model.Department, error) {
	logger.Infof("Start UpdateDepartment %+v ", updatedDepartment)
	var department model.Department
	if err := r.db.Where("id = ?", id).Preload("Department").First(&department).Error; err != nil {
		logger.Error("Error while fetching from department repo", err.Error())
		return model.Department{}, err
	}
	err := r.db.Model(&department).Updates(&updatedDepartment).Preload("Department").First(&department).Error
	logger.Info("End UpdateDepartment")
	return department, err
}

func (r *repository) UpdateDepartmentDetails(updatedDepartmentDetails model.DepartmentDetails, id string) (model.DepartmentDetails, error) {
	logger.Infof("Start UpdateDepartmentDetails %+v ", updatedDepartmentDetails)
	var departmentDetails model.DepartmentDetails
	if err := r.db.Where("id = ?", id).First(&departmentDetails).Error; err != nil {
		logger.Error("Error while fetching from department details repo", err.Error())
		return model.DepartmentDetails{}, err
	}
	err := r.db.Model(&departmentDetails).Updates(&updatedDepartmentDetails).Error
	logger.Info("End UpdateDepartmentDetails")
	return departmentDetails, err
}

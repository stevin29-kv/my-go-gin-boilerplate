package department

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/helpers"
	"employee-app/internal/entity/model"
	"employee-app/logger"
	"strconv"
)

func (ds *service) CreateDepartment(createDepartmentDto dto.CreateDepartment) *model.APIResponse {
	logger.Infof("Start CreateDepartment %+v", createDepartmentDto)

	tx := ds.db.Begin()

	departmentDetails := model.DepartmentDetails{
		DepartmentRoom: createDepartmentDto.DepartmentRoom,
		DepartmentCode: createDepartmentDto.DepartmentCode,
		Website:        createDepartmentDto.Website,
	}

	department := model.Department{
		Name:       createDepartmentDto.Name,
		Department: departmentDetails,
	}

	department, err := ds.department.CreateDepartment(department, tx)
	if err != nil {
		logger.Error("Error while creating department", err.Error())
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Unable to create department",
			},
		}
	}
	tx.Commit()

	logger.Infof("End CreateDepartment %+v", department)
	return &model.APIResponse{
		StatusCode: 201,
		Data:       department,
	}
}

func (ds *service) GetAllDepartments() *model.APIResponse {
	logger.Info("Start GetAllDepartments")
	departments, err := ds.department.GetAllDepartments()
	if err != nil {
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: helpers.ErrDepartmentNotFound.Error(),
			},
		}
	}
	logger.Infof("End GetAllDepartments count %d", len(departments))
	return &model.APIResponse{
		StatusCode: 200,
		Data:       departments,
	}
}

func (ds *service) GetDepartmentById(id string) *model.APIResponse {
	logger.Info("Start GetDepartmentById")
	department, err := ds.department.GetDepartmentById(id)
	if err != nil {
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: helpers.ErrDepartmentNotFound.Error(),
			},
		}
	}

	logger.Infof("End GetDepartmentById")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       department,
	}
}

func (ds *service) UpdateDepartment(UpdateDepartmentDto dto.UpdateDepartment, id string) *model.APIResponse {
	logger.Infof("Start UpdateDepartment %+v", UpdateDepartmentDto)

	tx := ds.db.Begin()

	department, err := ds.department.GetDepartmentById(id)
	if err != nil {
		logger.Error("Department not found", err.Error())
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Department not found",
			},
		}
	}
	deptDetailsId := department.DepartmentDetailsID

	updatedDepartmentDetails := model.DepartmentDetails{
		DepartmentRoom: UpdateDepartmentDto.DepartmentRoom,
		DepartmentCode: UpdateDepartmentDto.DepartmentCode,
		Website:        UpdateDepartmentDto.Website,
	}

	updatedDepartmentDetails, err = ds.department.UpdateDepartmentDetails(updatedDepartmentDetails, strconv.Itoa(deptDetailsId), tx)
	if err != nil {
		logger.Error("Error while updating department details", err.Error())
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Unable to update department details",
			},
		}
	}

	updatedDepartment := model.Department{
		Name:       UpdateDepartmentDto.Name,
		Department: updatedDepartmentDetails,
	}

	updatedDepartment, err = ds.department.UpdateDepartment(updatedDepartment, id, tx)
	if err != nil {
		logger.Error("Error while updating department", err.Error())
		tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Unable to update department",
			},
		}
	}
	tx.Commit()

	logger.Infof("End UpdateDepartment %+v", updatedDepartment)
	return &model.APIResponse{
		StatusCode: 200,
		Data:       updatedDepartment,
	}
}

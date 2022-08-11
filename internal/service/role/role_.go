package role

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/helpers"
	"employee-app/internal/entity/model"
	"employee-app/logger"
)

func (rs *service) CreateRole(createRoleDto dto.CreateRole) *model.APIResponse {
	logger.Infof("Start CreateRole %+v", createRoleDto)
	role := model.Role{
		Role: createRoleDto.Role,
	}

	// tx := rs.db.Begin()

	role, err := rs.role.CreateRole(role)
	if err != nil {
		logger.Error("Error while creating role", err.Error())
		// tx.Rollback()
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Unable to create role",
			},
		}
	}
	// tx.Commit()

	logger.Infof("End CreateRole %+v", role)
	return &model.APIResponse{
		StatusCode: 201,
		Data:       role,
	}
}

func (rs *service) GetAllRoles() *model.APIResponse {
	logger.Info("Start GetAllRoles")
	roles, err := rs.role.GetAllRoles()
	if err != nil {
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: helpers.ErrRoleNotFound.Error(),
			},
		}
	}
	logger.Infof("End GetAllRoles count %d", len(roles))
	return &model.APIResponse{
		StatusCode: 200,
		Data:       roles,
	}
}

func (rs *service) GetRoleById(id string) *model.APIResponse {
	logger.Info("Start GetRoleById")
	role, err := rs.role.GetRoleById(id)
	if err != nil {
		return &model.APIResponse{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: helpers.ErrRoleNotFound.Error(),
			},
		}
	}
	logger.Infof("End GetRoleById")
	return &model.APIResponse{
		StatusCode: 200,
		Data:       role,
	}
}

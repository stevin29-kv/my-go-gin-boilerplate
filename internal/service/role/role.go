package role

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/model"
	"employee-app/internal/repository/role"
)

type ServiceInterface interface {
	CreateRole(createRoleDto dto.CreateRole) *model.APIResponse
	GetAllRoles() *model.APIResponse
	GetRoleById(id string) *model.APIResponse
}

type service struct {
	role role.RepositoryInterface
}

func InitRoleService(role role.RepositoryInterface) ServiceInterface {
	return &service{
		role: role,
	}
}

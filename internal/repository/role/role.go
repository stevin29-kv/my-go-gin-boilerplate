package role

import (
	"employee-app/internal/entity/model"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CreateRole(role model.Role, tx *gorm.DB) (model.Role, error)
	GetAllRoles() ([]model.Role, error)
	GetRoleById(id string) (model.Role, error)
}

type repository struct {
	db *gorm.DB
}

func InitiRoleRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

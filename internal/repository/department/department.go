package department

import (
	"employee-app/internal/entity/model"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CreateDepartment(department model.Department, tx *gorm.DB) (model.Department, error)
	GetAllDepartments() ([]model.Department, error)
	GetDepartmentById(id string) (model.Department, error)
	UpdateDepartment(updatedDepartment model.Department, id string, tx *gorm.DB) (model.Department, error)
	UpdateDepartmentDetails(updatedDepartmentDetails model.DepartmentDetails, id string, tx *gorm.DB) (model.DepartmentDetails, error)
}

type repository struct {
	db *gorm.DB
}

func InitDepartmentRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

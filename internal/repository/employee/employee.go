package employee

import (
	"employee-app/internal/entity/helpers"
	"employee-app/internal/entity/model"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CreateEmployee(model.Employee) (model.Employee, error)
	GetAllEmployees(helpers.Pagination) ([]model.Employee, error)
	GetEmployeeById(string) (model.Employee, error)
	CreateAddress(model.Address) (model.Address, error)
	GetAddressById(int) (model.Address, error)
	DeleteEmployee(string) error
	UpdateEmployee(string, model.Employee) (model.Employee, error)
	UpdateAddress(string, model.Address) (model.Address, error)
	UpdateEmployeeStatusById(string, bool) (model.Employee, error)
	GetEmployeeByEmail(string) (model.Employee, error)
	UploadIdProof(id string, newFileName string) (model.Employee, error)
}

type repository struct {
	db *gorm.DB
}

func InitEmployeeRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

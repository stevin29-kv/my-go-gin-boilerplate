package user

import (
	"employee-app/internal/entity/model"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	CreateUser(model.User) (model.User, error)
	GetUserByEmail(string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) RepositoryInterface {
	return &repository{
		db: db,
	}
}

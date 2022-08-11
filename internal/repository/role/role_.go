package role

import (
	"employee-app/internal/entity/model"
	"employee-app/logger"

	"gorm.io/gorm"
)

func (r *repository) CreateRole(role model.Role, tx *gorm.DB) (model.Role, error) {
	logger.Infof("Start CreateRole %+v ", role)
	err := tx.Create(&role).Error
	logger.Info("End CreateRole")
	return role, err
}

func (r *repository) GetAllRoles() ([]model.Role, error) {
	var roles []model.Role
	res := r.db.Find(&roles)
	if res.Error != nil {
		msg := res.Error
		return nil, msg
	}
	return roles, res.Error
}

func (r *repository) GetRoleById(id string) (model.Role, error) {
	logger.Info("Start GetRoleById")
	var role model.Role
	response := r.db.Where("id =?", id).First(&role)
	logger.Info(response.Error)
	if response.Error != nil {
		logger.Error("Error while fetching from role repo", response.Error.Error())
	}
	logger.Infof("End GetRoleById")
	return role, response.Error
}

package user

import (
	"employee-app/internal/entity/model"
	"employee-app/logger"
)

func (r *repository) CreateUser(user model.User) (model.User, error) {
	logger.Info("Started CreateUser in Repo")
	err := r.db.Create(&user).Error
	logger.Info("Ended CreateUser in Repo")
	return user, err
}

func (r *repository) GetUserByEmail(email string) (model.User, error) {
	logger.Info("Started GetUserByEmail in Repo")
	var user model.User
	err := r.db.Find(&user, "email = ?", email).Error
	logger.Info(user)
	logger.Info("Ended GetUserByEmail in Repo")
	return user, err
}

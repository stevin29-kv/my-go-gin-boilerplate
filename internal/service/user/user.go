package user

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/model"
	"employee-app/internal/repository/user"
)

type ServiceInterface interface {
	CreateUser(dto.UserSignUpRequest) *model.APIResponseWithError
	UserLogin(dto.UserLoginRequest) *model.APIResponseWithError
}

type service struct {
	user user.RepositoryInterface
}

func InitUserService(user user.RepositoryInterface) ServiceInterface {
	return &service{
		user: user,
	}
}

package user

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/model"
	"employee-app/internal/utils"
	"employee-app/logger"
	"errors"
)

func (us *service) CreateUser(userData dto.UserSignUpRequest) *model.APIResponseWithError {
	if ok := utils.VerfityPassword(userData.Password, userData.ConfirmPassword); !ok {
		return &model.APIResponseWithError{
			StatusCode: 400,
			Data:       "Passwords doesn't match",
			Error:      nil,
		}
	}

	hashedPassword, ok := utils.HashPassword(userData.Password)
	if ok != nil {
		return &model.APIResponseWithError{
			StatusCode: 500,
			Data:       "Error while hashing password",
			Error:      nil,
		}
	}

	if ok := utils.ValidMailAddress(userData.Email); !ok {
		return &model.APIResponseWithError{
			StatusCode: 400,
			Data:       "Invalid Email",
			Error:      nil,
		}
	}

	user := model.User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Password:  hashedPassword,
	}
	logger.Info(user)

	// var returnData ReturnData
	user, err := us.user.CreateUser(user)
	if err != nil {
		logger.Error("Error while creating user")
		return &model.APIResponseWithError{
			StatusCode: 400,
			Data: &model.ErrorStatus{
				Message: "Cannot save user",
			},
			Error: err,
		}
	}
	// returnData.User = user
	// tokens, err := auth.GenerateAccessAndRefreshToken(user.Email)
	// if err != nil {
	// 	logger.Error("Error while creating tokens")
	// 	return &model.APIResponseWithError{
	// 		StatusCode: 400,
	// 		Data: &model.ErrorStatus{
	// 			Message: "Cannot generate token",
	// 		},
	// 		Error: err,
	// 	}
	// }
	// returnData.Token = tokens.(auth.TokenStruct)
	// logger.Info(returnData)
	logger.Info("Saved user")
	return &model.APIResponseWithError{
		StatusCode: 201,
		Data:       user,
		Error:      nil,
	}
}

func (us *service) UserLogin(loginData dto.UserLoginRequest) *model.APIResponseWithError {
	if ok := utils.ValidMailAddress(loginData.Email); !ok {
		return &model.APIResponseWithError{
			StatusCode: 400,
			Data:       "Invalid Email",
			Error:      errors.New("invalid email"),
		}
	}

	// var returnData ReturnData
	user, err := us.user.GetUserByEmail(loginData.Email)
	if err != nil {
		logger.Error("Error while getting user")
		return &model.APIResponseWithError{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Cannot get user",
			},
			Error: err,
		}
	}

	// returnData.User = user
	logger.Info(user)
	if ok := utils.CheckPasswordHash(loginData.Password, user.Password); !ok {
		return &model.APIResponseWithError{
			StatusCode: 404,
			Data: &model.ErrorStatus{
				Message: "Password doesn't match",
			},
			Error: errors.New("password not correct"),
		}
	}

	// tokens, err := auth.GenerateAccessAndRefreshToken(user.Email)
	// if err != nil {
	// 	logger.Error("Error while creating tokens")
	// 	return &model.APIResponseWithError{
	// 		StatusCode: 400,
	// 		Data: &model.ErrorStatus{
	// 			Message: "Cannot generate token",
	// 		},
	// 		Error: err,
	// 	}
	// }

	// returnData.Token = tokens.(auth.TokenStruct)
	logger.Info("Login user")
	return &model.APIResponseWithError{
		StatusCode: 200,
		Data:       user,
		Error:      nil,
	}
}

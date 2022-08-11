package controller

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/service/user"
	"employee-app/logger"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.ServiceInterface
}

func InitUserController(us user.ServiceInterface) *userController {
	return &userController{
		userService: us,
	}
}

func (uc *userController) CreateUser(c *gin.Context) {
	var userData dto.UserSignUpRequest
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(400, "Error while binding")
		return
	}
	resp := uc.userService.CreateUser(userData)
	logger.Info(resp.Data)
	if resp.Error != nil {
		c.JSON(resp.StatusCode, resp.Data)
		return
	}

	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("access", resp.Data.(service.ReturnData).Token.Access, 60*60*24, "/", "http://localhost:8080", true, true)
	// c.SetCookie("refresh", resp.Data.(service.ReturnData).Token.Refresh, 60*60*24, "/", "http://localhost:8080", true, true)

	logger.Info("Successfully Signed Up")
	c.JSON(resp.StatusCode, resp.Data)
}

func (uc *userController) LoginUser(c *gin.Context) {
	var userData dto.UserLoginRequest
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(400, "Error while binding")
		return
	}
	resp := uc.userService.UserLogin(userData)
	logger.Info(resp.Data)
	if resp.Error != nil {
		c.JSON(resp.StatusCode, resp.Data)
		return
	}

	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("access", resp.Data.(service.ReturnData).Token.Access, 60*60*24, "/", "http://localhost:8080", true, true)
	// c.SetCookie("refresh", resp.Data.(service.ReturnData).Token.Refresh, 60*60*24, "/", "http://localhost:8080", true, true)

	logger.Info("Successfully Logged In")

	c.JSON(resp.StatusCode, resp.Data)
}

func (uc *userController) LogoutUser(c *gin.Context) {
	// c.SetSameSite(http.SameSiteNoneMode)
	// c.SetCookie("access", "", -1, "/", "http://localhost:8080", true, true)
	// c.SetCookie("refresh", "", -1, "/", "http://localhost:8080", true, true)
	c.JSON(200, "Successfully Logged Out")
}

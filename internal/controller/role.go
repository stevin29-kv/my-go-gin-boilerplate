package controller

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/service/role"
	"employee-app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	roleService role.ServiceInterface
}

func InitRoleController(rs role.ServiceInterface) *roleController {
	return &roleController{
		roleService: rs,
	}
}

func (rc *roleController) CreateRole(c *gin.Context) {
	logger.Info("Start CreateRole in Controller")
	var createRoleDto dto.CreateRole
	if err := c.BindJSON(&createRoleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := rc.roleService.CreateRole(createRoleDto)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End CreateRole in Controller")
}

func (rc *roleController) GetAllRoles(c *gin.Context) {
	logger.Info("Start getAllRoles in Controller")
	resp := rc.roleService.GetAllRoles()
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End getAllRoles in Controller")
}

func (rc *roleController) GetRoleById(c *gin.Context) {
	id := c.Param("id")
	logger.Info("Start GetRoleById in Controller")
	resp := rc.roleService.GetRoleById(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetRoleById in Controller")
}

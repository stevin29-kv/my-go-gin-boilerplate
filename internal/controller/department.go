package controller

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/service/department"
	"employee-app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type departmentController struct {
	departmentService department.ServiceInterface
}

func InitDepartmentController(ds department.ServiceInterface) *departmentController {
	return &departmentController{
		departmentService: ds,
	}
}

func (dc *departmentController) CreateDepartment(c *gin.Context) {
	logger.Info("Start CreateDepartment in Controller")
	var createDepartmentDto dto.CreateDepartment
	if err := c.BindJSON(&createDepartmentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := dc.departmentService.CreateDepartment(createDepartmentDto)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End CreateDepartment in Controller")
}

func (dc *departmentController) GetAllDepartments(c *gin.Context) {
	logger.Info("Start GetAllDepartments in Controller")
	resp := dc.departmentService.GetAllDepartments()
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetAllDepartments in Controller")
}

func (dc *departmentController) GetDepartmentById(c *gin.Context) {
	id := c.Param("id")
	logger.Info("Start GetDepartmentById in Controller")
	resp := dc.departmentService.GetDepartmentById(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetDepartmentById in Controller")
}

func (dc *departmentController) UpdateDepartment(c *gin.Context) {
	logger.Info("Start UpdateDepartment in Controller")
	var UpdateDepartmentDto dto.UpdateDepartment
	if err := c.BindJSON(&UpdateDepartmentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	resp := dc.departmentService.UpdateDepartment(UpdateDepartmentDto, id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UpdateDepartment in Controller")
}

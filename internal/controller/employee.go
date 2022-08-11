package controller

import (
	"employee-app/internal/entity/dto"
	"employee-app/internal/entity/helpers"
	"employee-app/internal/service/employee"
	"employee-app/logger"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type employeeController struct {
	employeeService employee.ServiceInterface
}

func InitEmployeeController(ue employee.ServiceInterface) *employeeController {
	return &employeeController{
		employeeService: ue,
	}
}

func (ec *employeeController) CreateEmployee(c *gin.Context) {
	logger.Info("Start CreateEmployee in Controller")
	var employeeData dto.CreateEmployeeRequest
	if err := c.BindJSON(&employeeData); err != nil {
		logger.Error(err)
		c.JSON(400, helpers.ErrInvalidRequest)
		return
	}

	resp := ec.employeeService.CreateEmployee(employeeData)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End CreateEmployee in Controller")
}

func (ec *employeeController) GetAllEmployees(c *gin.Context) {
	logger.Info("Start GetAllEmployees in Controller")
	search, err := helpers.GetFilterValue(c)
	if err != nil {
		logger.Info("search value is empty")
	}
	sort_by, order, err := helpers.GetSortingValue(c)
	if err != nil {
		logger.Info("sort value is empty")
	}
	resp := ec.employeeService.GetAllEmployees(search, sort_by, order)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetAllEmployees in Controller")
}

func (ec *employeeController) GetEmployeeById(c *gin.Context) {
	logger.Info("Start GetEmployeeById - Controller")
	id := c.Param("id")
	resp := ec.employeeService.GetEmployeeById(id)
	logger.Info(resp.Data)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End GetEmployeeById - Controller")
}

func (ec *employeeController) DeleteEmployee(c *gin.Context) {
	logger.Info("Start DeleteEmployee - Controller")
	id := c.Param("id")
	resp := ec.employeeService.DeleteEmployee(id)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End DeleteEmployee - Controller")
}

func (ec *employeeController) UpdateEmployee(c *gin.Context) {
	logger.Info("Start UpdateEmployee - COntroller")
	id := c.Param("id")
	var employeeData dto.UpdateEmployeeRequest
	if err := c.BindJSON(&employeeData); err != nil {
		c.JSON(400, helpers.ErrInvalidRequest)
		return
	}
	resp := ec.employeeService.UpdateEmployee(id, employeeData)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UpdateEmployee - Controller")
}

func (ec *employeeController) UpdateEmployeeStatusById(c *gin.Context) {
	logger.Info("Start UpdateEmployeeStatusById - COntroller")
	id := c.Param("id")
	var employeeData dto.UpdateEmployeeStatusRequest
	if err := c.BindJSON(&employeeData); err != nil {
		c.JSON(400, helpers.ErrInvalidRequest)
		return
	}
	resp := ec.employeeService.UpdateEmployeeStatusById(id, employeeData)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UpdateEmployeeStatusById - Controller")
}

func (ec *employeeController) UploadIdProof(c *gin.Context) {
	logger.Info("Start UploadIdProof in Controller")
	id := c.Param("id")
	file, err := c.FormFile("file")

	if err != nil {
		logger.Error(err)
		// logger.Error("No file")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	logger.Info(file.Filename)

	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "./assets/"+newFileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		logger.Info(err)
		return
	}

	resp := ec.employeeService.UploadIdProof(id, newFileName)
	c.JSON(resp.StatusCode, resp.Data)
	logger.Info("End UploadIdProof in Controller")
}

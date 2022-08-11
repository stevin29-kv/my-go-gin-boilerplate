package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthController struct{}

func InitHealthController() *healthController {
	return &healthController{}
}

func (h *healthController) HandleGetHealth(c *gin.Context) {
	message := "Health is UP"
	c.JSON(http.StatusOK, message)
}

package router

import "github.com/gin-gonic/gin"

func (c *CapsuleRouter) HealthRoutes(r *gin.RouterGroup) {
	healthController := c.Controller.HealthController

	r.GET("/health", healthController.HandleGetHealth)
}

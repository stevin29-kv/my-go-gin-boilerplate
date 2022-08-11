package router

import (
	"employee-app/internal/controller"
	"employee-app/internal/repository"
	"employee-app/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CapsuleRouter struct {
	DB         *gorm.DB
	Repository *repository.Repository
	Service    *service.Service
	Controller *controller.Controller
}

func PrepareRouter(capsule *CapsuleRouter) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/employee-app/health"),
		gin.Recovery(),
	)

	engine := router.Group("employee-app")
	capsule.HealthRoutes(engine)

	v1 := engine.Group("v1")
	capsule.EmployeeRouter(v1)
	capsule.DepartmentRouter(v1)
	capsule.RoleRouter(v1)
	capsule.UserRouter(v1)

	return router
}

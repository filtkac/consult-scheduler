package config

import (
	"consult-scheduler/controller"
	"consult-scheduler/model"
	"consult-scheduler/repository"
	"consult-scheduler/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RouterConfig(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(ErrorHandler())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.ErrorResponse{Message: "Path does not exist."})
	})

	consultTypeController := initConsultTypeController(db)
	consultTypeRouter := router.Group("/consult-types")
	consultTypeRouter.GET("", consultTypeController.GetAllConsultTypes)
	consultTypeRouter.DELETE("/:id", consultTypeController.DeleteConsultType)
	consultTypeRouter.POST("", consultTypeController.CreateConsultType)

	departmentController := initDepartmentController(db)
	departmentRouter := router.Group("/departments")
	departmentRouter.GET("", departmentController.GetAllDepartments)
	departmentRouter.DELETE("/:id", departmentController.DeleteDepartment)
	departmentRouter.POST("", departmentController.CreateDepartment)

	dayTemplateController := initDayTemplateController(db)
	dayTemplateRouter := router.Group("/day-templates")
	dayTemplateRouter.GET("", dayTemplateController.GetDayTemplates)
	dayTemplateRouter.DELETE("/:id", dayTemplateController.DeleteDayTemplate)
	dayTemplateRouter.POST("", dayTemplateController.CreateDayTemplate)
	dayTemplateRouter.GET(":id/consults", dayTemplateController.GetConsultsForDayTemplate)
	dayTemplateRouter.POST(":id/consults", dayTemplateController.CreateConsultsForDayTemplate)

	return router
}

func initConsultTypeController(db *gorm.DB) *controller.ConsultTypeController {
	consultTypeRepository := repository.NewConsultTypeRepository(db)
	consultTypeService := service.NewConsultTypeService(consultTypeRepository)
	return controller.NewConsultTypeController(consultTypeService)
}

func initDepartmentController(db *gorm.DB) *controller.DepartmentController {
	departmentRepository := repository.NewDepartmentRepository(db)
	departmentService := service.NewDepartmentService(departmentRepository)
	return controller.NewDepartmentController(departmentService)
}

func initDayTemplateController(db *gorm.DB) *controller.DayTemplateController {
	dayTemplateRepository := repository.NewDayTemplateRepository(db)
	dayTemplateService := service.NewDayTemplateService(dayTemplateRepository)
	return controller.NewDayTemplateController(dayTemplateService)
}

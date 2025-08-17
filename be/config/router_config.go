package config

import (
	"github.com/filtkac/consult-scheduler/controller"
	"github.com/filtkac/consult-scheduler/model"
	"github.com/filtkac/consult-scheduler/repository"
	"github.com/filtkac/consult-scheduler/service"
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
	consultTypeRouter.DELETE("/:consultTypeId", consultTypeController.DeleteConsultType)
	consultTypeRouter.POST("", consultTypeController.CreateConsultType)

	departmentController := initDepartmentController(db)
	departmentRouter := router.Group("/departments")
	departmentRouter.GET("", departmentController.GetAllDepartments)
	departmentRouter.DELETE("/:departmentId", departmentController.DeleteDepartment)
	departmentRouter.POST("", departmentController.CreateDepartment)

	dayTemplateController := initDayTemplateController(db)
	dayTemplateRouter := router.Group("/day-templates")
	dayTemplateRouter.GET("", dayTemplateController.GetDayTemplates)
	dayTemplateRouter.DELETE("/:dayTemplateId", dayTemplateController.DeleteDayTemplate)
	dayTemplateRouter.POST("", dayTemplateController.CreateDayTemplate)
	dayTemplateRouter.GET("/:dayTemplateId/consults", dayTemplateController.GetConsultsForDayTemplate)
	dayTemplateRouter.POST("/:dayTemplateId/consults", dayTemplateController.CreateConsultsForDayTemplate)

	consultController := initConsultController(db)
	consultRouter := router.Group("/departments/:departmentId/consults")
	consultRouter.GET("", consultController.GetDepartmentConsults)
	consultRouter.POST("", consultController.CreateDepartmentConsult)
	consultRouter.DELETE("/:consultId", consultController.DeleteDepartmentConsult)
	consultRouter.POST("/:consultId", consultController.UpdateDepartmentConsult)

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

func initConsultController(db *gorm.DB) *controller.ConsultController {
	consultRepository := repository.NewConsultRepository(db)
	consultService := service.NewConsultService(consultRepository)
	return controller.NewConsultController(consultService)
}

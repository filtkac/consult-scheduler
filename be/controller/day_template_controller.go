package controller

import (
	"consult-scheduler/model"
	"consult-scheduler/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DayTemplateController struct {
	s *service.DayTemplateService
}

func NewDayTemplateController(s *service.DayTemplateService) *DayTemplateController {
	return &DayTemplateController{s: s}
}

func (c DayTemplateController) GetDayTemplates(ctx *gin.Context) {
	departmentIDParam := ctx.Query("department-id")
	if departmentIDParam == "" {
		result, err := c.s.GetAllDayTemplates(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	} else {
		departmentID, err := strconv.ParseUint(departmentIDParam, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusBadRequest,
				model.ErrorResponse{Message: "Department ID must be a positive number."})
			return
		}
		result, err := c.s.GetDayTemplatesForDepartment(ctx, uint(departmentID))
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func (c DayTemplateController) DeleteDayTemplate(ctx *gin.Context) {
	id, err := GetIDPathParamOrAbort(ctx)
	if err != nil {
		return
	}
	err = c.s.DeleteDayTemplate(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c DayTemplateController) CreateDayTemplate(ctx *gin.Context) {
	var dayTemplate model.DayTemplateDto
	if err := ctx.ShouldBind(&dayTemplate); err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := c.s.CreateDayTemplate(ctx, &dayTemplate)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func (c DayTemplateController) GetConsultsForDayTemplate(ctx *gin.Context) {
	dayTemplateID, err := GetIDPathParamOrAbort(ctx)
	if err != nil {
		return
	}
	result, err := c.s.GetConsultsForDayTemplate(ctx, dayTemplateID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c DayTemplateController) CreateConsultsForDayTemplate(ctx *gin.Context) {
	dayTemplateID, err := GetIDPathParamOrAbort(ctx)
	if err != nil {
		return
	}
	var consults []*model.DayTemplateConsultDto
	if err := ctx.ShouldBind(&consults); err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := c.s.CreateConsultsForDayTemplate(ctx, dayTemplateID, consults)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

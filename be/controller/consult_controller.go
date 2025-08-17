package controller

import (
	"fmt"
	"github.com/filtkac/consult-scheduler/model"
	"github.com/filtkac/consult-scheduler/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ConsultController struct {
	s *service.ConsultService
}

func NewConsultController(s *service.ConsultService) *ConsultController {
	return &ConsultController{s: s}
}

func (c ConsultController) GetDepartmentConsults(ctx *gin.Context) {
	departmentID, err := GetIDPathParamOrAbort(ctx, "departmentId")
	if err != nil {
		return
	}

	dayParam := ctx.Query("day")
	if dayParam == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "Query parameter 'day' must be specified"})
		return
	}
	day, err := time.Parse("2006-01-02", dayParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Message: fmt.Sprintf("Failed to parse day: %s, error: %s", dayParam, err.Error())})
		return
	}

	consults, err := c.s.GetDepartmentConsultsForDay(ctx, departmentID, day)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, consults)
}

func (c ConsultController) DeleteDepartmentConsult(ctx *gin.Context) {
	departmentId, err := GetIDPathParamOrAbort(ctx, "departmentId")
	if err != nil {
		return
	}

	consultId, err := GetIDPathParamOrAbort(ctx, "consultId")
	if err != nil {
		return
	}

	err = c.s.DeleteDepartmentConsult(ctx, departmentId, consultId)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c ConsultController) CreateDepartmentConsult(ctx *gin.Context) {
	departmentID, err := GetIDPathParamOrAbort(ctx, "departmentId")
	if err != nil {
		return
	}

	var consult model.CreateConsultDto
	if err := ctx.ShouldBind(&consult); err != nil {
		_ = ctx.Error(err)
		return
	}

	result, err := c.s.CreateDepartmentConsult(ctx, departmentID, &consult)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c ConsultController) UpdateDepartmentConsult(ctx *gin.Context) {
	departmentID, err := GetIDPathParamOrAbort(ctx, "departmentId")
	if err != nil {
		return
	}

	consultID, err := GetIDPathParamOrAbort(ctx, "consultId")
	if err != nil {
		return
	}

	var consult model.CreateConsultDto
	if err := ctx.ShouldBind(&consult); err != nil {
		_ = ctx.Error(err)
		return
	}

	result, err := c.s.UpdateDepartmentConsult(ctx, departmentID, consultID, &consult)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

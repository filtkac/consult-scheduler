package controller

import (
	"consult-scheduler/model"
	"consult-scheduler/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConsultTypeController struct {
	s *service.ConsultTypeService
}

func NewConsultTypeController(s *service.ConsultTypeService) *ConsultTypeController {
	return &ConsultTypeController{s: s}
}

func (c ConsultTypeController) GetAllConsultTypes(ctx *gin.Context) {
	consultTypes, err := c.s.GetAllConsultTypes(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, consultTypes)
}

func (c ConsultTypeController) DeleteConsultType(ctx *gin.Context) {
	id, err := GetIDPathParamOrAbort(ctx, "consultTypeId")
	if err != nil {
		return
	}
	err = c.s.DeleteConsultType(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c ConsultTypeController) CreateConsultType(ctx *gin.Context) {
	var consultType model.ConsultTypeDto
	if err := ctx.ShouldBind(&consultType); err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := c.s.CreateConsultType(ctx, &consultType)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

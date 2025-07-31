package controller

import (
	"consult-scheduler/model"
	"consult-scheduler/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DepartmentController struct {
	s *service.DepartmentService
}

func NewDepartmentController(s *service.DepartmentService) *DepartmentController {
	return &DepartmentController{s: s}
}

func (c DepartmentController) GetAllDepartments(ctx *gin.Context) {
	result, err := c.s.GetAllDepartments(ctx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c DepartmentController) DeleteDepartment(ctx *gin.Context) {
	id, err := GetIDPathParamOrAbort(ctx)
	if err != nil {
		return
	}
	err = c.s.DeleteDepartment(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c DepartmentController) CreateDepartment(ctx *gin.Context) {
	var department model.DepartmentDto
	if err := ctx.ShouldBind(&department); err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := c.s.CreateDepartment(ctx, &department)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

package controller

import (
	"consult-scheduler/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIDPathParamOrAbort(ctx *gin.Context) (uint, error) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Message: "ID must be a positive number."})
		return 0, err
	}
	return uint(id), nil
}

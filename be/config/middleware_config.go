package config

import (
	"consult-scheduler/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"net/http"
)

// ErrorHandler Adopted from https://gin-gonic.com/en/docs/examples/error-handling-middleware/
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			var validationError validator.ValidationErrors
			if errors.As(err, &validationError) {
				validationErrors := make([]string, 0)
				for _, err := range validationError {
					validationErrors = append(validationErrors, err.Error())
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ValidationErrorResponse{ValidationErrors: validationErrors})
				return
			}

			var customValidationError model.CustomValidationError
			if errors.As(err, &customValidationError) {
				c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Message: customValidationError.Error()})
				return
			}

			var pgError *pgconn.PgError
			if errors.As(err, &pgError) {
				// 23505 violates unique constraint
				// 23503 violates foreign key constraint
				if pgError.Code == "23505" || pgError.Code == "23503" {
					c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{Message: pgError.Detail})
					return
				}

				// 23502 violates not-null constraint
				if pgError.Code == "23502" {
					c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
						Message: fmt.Sprintf("'%v' cannot be null", pgError.ColumnName),
					})
					return
				}
			}

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		}
	}
}

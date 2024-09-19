package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tinwritescode/gin-tin/pkg/model"
)

func HandleValidationErrors(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]model.ValidationError, len(ve))
		for i, fe := range ve {
			out[i] = model.ValidationError{
				Field:   strings.ToLower(fe.Field()),
				Message: getErrorMsg(fe),
			}
		}
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error:   "Validation failed",
			Details: out,
		})
	} else {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: err.Error(),
		})
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "min":
		return fmt.Sprintf("Should be at least %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Should be at most %s characters long", fe.Param())
	default:
		return "Invalid value"
	}
}

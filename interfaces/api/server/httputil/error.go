package httputil

import (
	"github.com/gin-gonic/gin"
)

// HTTPError is response error structure
type HTTPError struct {
	Code    int    `json:"Code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewHTTPError sets error response to ctx
func NewHTTPError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

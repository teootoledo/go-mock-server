package controller

import (
	"mock-server/internal/external/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetErrorResponseBadRequest(ctx *gin.Context, code int, message, details string) {
	ctx.JSON(code, resources.ErrorResponseHTTPBadRequest{
		StatusCode: http.StatusBadRequest,
		ErrorResponse: resources.ErrorResponse{
			Message: message,
			Details: details,
		},
	})
}

func SetErrorResponseNotFound(ctx *gin.Context, code int, message, details string) {
	ctx.JSON(code, resources.ErrorResponseHTTPNotFound{
		StatusCode: http.StatusNotFound,
		ErrorResponse: resources.ErrorResponse{
			Message: message,
			Details: details,
		},
	})
}

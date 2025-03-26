package common

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var ErrResourceNotFound = errors.New("resource not found")
var ErrInternalServer = errors.New("internal server error")
var ErrBadRequest = errors.New("bad request")
var ErrUnauthorized = errors.New("unauthorized")
var ErrConflict = errors.New("conflict")

func HandleError(ctx *gin.Context, err error) {
	var statusCode int
	var message string
	var stacktrace string
	shouldShowStacktrace := true

	if os.Getenv("MODE") == "release" {
		shouldShowStacktrace = false
	}

	switch err {
	case ErrResourceNotFound:
		statusCode = http.StatusNotFound
		message = "Resource not found"
	case ErrInternalServer:
		statusCode = http.StatusInternalServerError
		message = "Internal server error"
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
		message = "Bad request"
	case ErrUnauthorized:
		statusCode = http.StatusUnauthorized
		message = "Unauthorized"
	case ErrConflict:
		statusCode = http.StatusConflict
		message = "Business rule conflict"
	default:
		statusCode = http.StatusInternalServerError
		message = "Internal server error"
	}

	if shouldShowStacktrace {
		stacktrace = err.Error()
	}

	ctx.JSON(statusCode, ErrorResponse{
		Message:    message,
		Stacktrace: stacktrace,
	})
}

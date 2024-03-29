package sprinter

import (
	"fmt"
	"log"
	"runtime"
)

const (
	StatusOK = 200
	// BadRequestErr request is invalid
	BadRequestErr = 400
	// UnauthorizedErr has not authorized
	UnauthorizedErr = 401
	// ForbiddenErr permission is invalid
	ForbiddenErr = 403
	// NotFoundError object not found
	NotFoundError = 404
	// InternalServerError server stopped
	InternalServerError = 500
)

func OnResponseError(code int, err error, message string) Response {
	if _, file, line, ok := runtime.Caller(1); ok {
		log.Println(fmt.Sprintf("%s:%d %s:%s", file, line, err, message))
	}
	return Response{
		Code: code,
		Text: message,
		Err:  err,
	}
}

package http_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	badRequest        = "BAD_REQUEST"
	statusOk          = "SUCCESS"
	statusServerError = "SERVER_ERROR"
)

type ResponseEntity struct {
	Status   int         `json:"status"`
	Response interface{} `json:"response"`
}

func getResponse(status int, h interface{}) *ResponseEntity {
	return &ResponseEntity{Status: status, Response: h}
}

func NewBadRequestResponse(message string) *ResponseEntity {
	return getResponse(http.StatusBadRequest, gin.H{"code": badRequest, "reason": message})
}

func NewOkResponse(message string) *ResponseEntity {
	return getResponse(http.StatusOK, gin.H{"code": statusOk, "message": message})
}

func NewInternalServerError(message string, err error) *ResponseEntity {
	return getResponse(http.StatusInternalServerError, gin.H{"code": statusServerError, "Message": message, "Error": err})
}

func NewServiceResponse(statusCode int, data interface{}) *ResponseEntity {
	return getResponse(statusCode, data)
}

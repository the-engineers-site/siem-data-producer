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

type Response struct {
	Status  int   `json:"status"`
	Message gin.H `json:"message"`
}

func getResponse(status int, h gin.H) *Response {
	return &Response{Status: status, Message: h}
}

func NewBadRequestResponse(message string) *Response {
	return getResponse(http.StatusBadRequest, gin.H{"code": badRequest, "reason": message})
}

func NewOkResponse(message string) *Response {
	return getResponse(http.StatusOK, gin.H{"code": statusOk, "message": message})
}

func NewInternalServerError(message string, err error) *Response {
	return getResponse(http.StatusInternalServerError, gin.H{"code": statusServerError, "Message": message, "Error": err})
}

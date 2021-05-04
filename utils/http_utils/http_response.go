package http_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	notFound   = "NOT_FOUND"
	badRequest = "BAD_REQUEST"
	statusOk   = "SUCCESS"
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

func NewNotFountResponse(message string) *Response {
	return getResponse(http.StatusNotFound, gin.H{"code": notFound, "reason": message})
}

func NewOkResponse(message string) *Response {
	return getResponse(http.StatusOK, gin.H{"code": statusOk, "message": message})
}

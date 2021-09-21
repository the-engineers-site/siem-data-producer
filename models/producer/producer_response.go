package producer

import "github.com/gin-gonic/gin"

type ResponseInterface interface {
	SetMessage(status int, message interface{}, err error)
	GetResponse() gin.H
	GetStatus() int
}

type Response struct {
	Status  int         `json:"status,omitempty"`
	Message interface{} `json:"response,omitempty"`
}

func (c *Response) GetStatus() int {
	return c.Status
}

func (c *Response) SetMessage(status int, message interface{}, err interface{}) {
	if err != nil {
		c.Message = gin.H{
			"Error": err,
		}
	} else {
		c.Message = message
	}
	c.Status = status
}

func (c *Response) GetResponse() interface{} {
	return c.Message
}

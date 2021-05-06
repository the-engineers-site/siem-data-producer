package health

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/yjagdale/siem-data-producer/services"
)

func Ping(c *gin.Context) {
	resp := services.HealthService.HealthCheck()
	c.JSON(resp.Status, resp.Response)
	return
}

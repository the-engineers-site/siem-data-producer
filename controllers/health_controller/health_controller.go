package health_controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/yjagdale/siem-data-producer/services"
)

func Ping(c *gin.Context) {
	resp := services.HealthService.Check()
	c.JSON(resp.Status, resp.Response)
	return
}

func Health(c *gin.Context) {
	resp := services.HealthService.Check()
	c.JSON(resp.Status, gin.H{"status": "ok"})
}

package configuration_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/models/configuration"
	"gitlab.com/yjagdale/siem-data-producer/services"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"strconv"
)

func SaveConfiguration(c *gin.Context) {
	var resp *http_utils.ResponseEntity
	var config configuration.Configuration

	err := c.ShouldBindJSON(&config)
	if err != nil {
		resp = http_utils.NewBadRequestResponse("Invalid body")
	} else {
		log.Infoln("Storing configuration")
		resp = services.ConfigurationService.SaveConfig(&config)
	}
	c.JSON(resp.Status, resp)
	return
}

func UpdateConfiguration(c *gin.Context) {
	var resp *http_utils.ResponseEntity
	var config configuration.Configuration

	err := c.ShouldBindJSON(&config)
	if err != nil {
		resp = http_utils.NewOkResponse("Invalid body")
	} else {
		log.Infoln("Storing configuration")
		resp = services.ConfigurationService.UpdateConfig(&config)
	}
	c.JSON(resp.Status, resp)
	return
}

func DeleteConfiguration(c *gin.Context) {
	var resp *http_utils.ResponseEntity
	var config configuration.Configuration

	err := c.ShouldBindJSON(&config)
	if err != nil {
		resp = http_utils.NewOkResponse("Invalid body")
	} else {
		log.Infoln("Storing configuration")
		resp = services.ConfigurationService.DeleteConfig(&config)
	}
	c.JSON(resp.Status, resp)
	return
}

func GetConfiguration(c *gin.Context) {
	var resp *http_utils.ResponseEntity
	log.Infoln("Fetching configuration")
	id := c.Param("id")
	if id == "" {
		resp = services.ConfigurationService.GetConfig(&configuration.Configuration{})
	} else {
		config := configuration.Configuration{}
		configId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			resp = http_utils.NewBadRequestResponse("Invalid ID")
		} else {
			config.ID = uint(configId)
			resp = services.ConfigurationService.GetConfig(&config)
		}
	}
	c.JSON(resp.Status, resp)
	return
}

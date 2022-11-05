package configuration_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"siem-data-producer/constants"
	"siem-data-producer/models/configuration"
	"siem-data-producer/services"
	"siem-data-producer/utils/http_utils"
	"strconv"
)

func SaveConfiguration(c *gin.Context) {
	var resp = configuration.Response{}
	var config configuration.Configuration

	err := c.ShouldBindJSON(&config)
	if err != nil {
		resp.SetMessage(http.StatusBadRequest, "Invalid body", err)
	} else {
		log.Infoln("Storing configuration")
		resp = services.ConfigurationService.SaveConfig(&config)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func ReloadConfiguration(c *gin.Context) {
	resp := services.ConfigurationService.Reload()
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func UpdateConfiguration(c *gin.Context) {
	var resp = configuration.Response{}
	var config configuration.Configuration

	err := c.ShouldBindJSON(&config)
	if err != nil {
		resp.SetMessage(http.StatusBadRequest, "Invalid Body", err)
	} else {
		log.Infoln("Storing configuration")
		resp = services.ConfigurationService.UpdateConfig(&config)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func DeleteConfiguration(c *gin.Context) {
	var resp = &configuration.Response{}
	var config []int
	id := c.Param("id")

	if id != "" {
		log.Infoln("Deleting configuration with object", id)
		intId, err := strconv.Atoi(id)
		if err != nil {
			resp.SetMessage(http.StatusBadRequest, "Invalid configuration ID", err)
			c.JSON(resp.GetStatus(), resp.GetResponse())
			return
		}
		config = append(config, intId)
	} else {
		err := c.ShouldBindJSON(&config)
		if err != nil {
			c.JSON(http.StatusBadRequest, http_utils.NewOkResponse(constants.ResponseBadRequest+err.Error()))
			return
		}
		log.Infoln("Deleting object", config)

	}
	resp = services.ConfigurationService.DeleteConfig(config)
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func GetOverrides(c *gin.Context) {
	resp := configuration.Response{}
	resp.SetMessage(http.StatusOK, constants.Executors, nil)
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func GetConfiguration(c *gin.Context) {
	var resp = configuration.Response{}
	log.Infoln("Fetching configuration")
	id := c.Param("id")
	if id == "" {
		resp = services.ConfigurationService.GetConfig(&configuration.Configuration{})
	} else {
		config := configuration.Configuration{}
		configId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			resp.SetMessage(http.StatusBadRequest, "Invalid Configuration ID", err)
		} else {
			config.ID = uint(configId)
			resp = services.ConfigurationService.GetConfig(&config)
		}
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

package services

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"siem-data-producer/constants"
	"siem-data-producer/models/configuration"
)

var (
	ConfigurationService ConfigurationServiceInterface = &configurationService{}
)

type ConfigurationServiceInterface interface {
	SaveConfig(configuration *configuration.Configuration) configuration.Response
	UpdateConfig(configuration *configuration.Configuration) configuration.Response
	DeleteConfig(configuration []int) *configuration.Response
	GetConfig(configuration *configuration.Configuration) configuration.Response
	Reload() configuration.Response
}

type configurationService struct {
}

func (c *configurationService) Reload() configuration.Response {
	var config configuration.Configuration
	resp := config.GetAll()

	if resp.GetStatus() != 200 {
		return resp
	}
	if constants.Executors == nil {
		constants.Executors = make(map[string][]string)
	}
	for k, v := range resp.GetResponse().([]configuration.Configuration) {
		log.Debugln("Loading", k, v.OverrideKey)
		constants.Executors[v.OverrideKey] = v.OverrideValues
	}
	resp = configuration.Response{}
	resp.SetMessage(http.StatusOK, gin.H{"message": "Reloaded successfully"}, nil)
	return resp
}

func (c *configurationService) UpdateConfig(configuration *configuration.Configuration) configuration.Response {
	return configuration.Update()
}

func (c *configurationService) DeleteConfig(ids []int) *configuration.Response {
	var success []int
	var failed []int
	var notFount []int
	var resp = configuration.Response{}
	for _, objectId := range ids {
		conf := configuration.Configuration{}
		conf.ID = uint(objectId)
		response := conf.Delete()
		if response.GetStatus() == 404 {
			notFount = append(notFount, objectId)
		} else if response.Status != 200 {
			failed = append(failed, objectId)
		} else {
			success = append(success, objectId)
		}
	}
	resp.SetMessage(http.StatusOK, gin.H{"Success": success, "Failed": failed, "NotFound": notFount}, nil)
	return &resp
}

func (c *configurationService) GetConfig(configuration *configuration.Configuration) configuration.Response {
	if configuration.ID == 0 {
		return configuration.GetAll()
	}
	return configuration.Get()
}

func (c configurationService) SaveConfig(config *configuration.Configuration) configuration.Response {
	var resp configuration.Response
	if config.Validate() != nil {
		log.Debugln("Validation failed for configuration.")
		err := config.Validate()
		resp.SetMessage(http.StatusBadRequest, "validation failed", err)

	}
	return config.Save()
}

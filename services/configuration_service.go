package services

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/models/configuration"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
)

var (
	ConfigurationService ConfigurationServiceInterface = &configurationService{}
)

type ConfigurationServiceInterface interface {
	SaveConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
	UpdateConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
	DeleteConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
	GetConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
}

type configurationService struct {
}

func (c *configurationService) UpdateConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity {
	return configuration.Update()
}

func (c *configurationService) DeleteConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity {
	return configuration.Delete()
}

func (c *configurationService) GetConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity {
	if configuration.ID == 0 {
		return configuration.GetAll()
	}
	return configuration.Get()
}

func (c configurationService) SaveConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity {
	if configuration.Validate() != nil {
		log.Debugln("Validation failed for configuration.")
		return configuration.Validate()
	}
	return configuration.Save()
}

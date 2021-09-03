package services

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/models/configuration"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"net/http"
)

var (
	ConfigurationService ConfigurationServiceInterface = &configurationService{}
)

type ConfigurationServiceInterface interface {
	SaveConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
	UpdateConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
	DeleteConfig(configuration []int) *http_utils.ResponseEntity
	GetConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity
}

type configurationService struct {
}

func (c *configurationService) UpdateConfig(configuration *configuration.Configuration) *http_utils.ResponseEntity {
	return configuration.Update()
}

func (c *configurationService) DeleteConfig(ids []int) *http_utils.ResponseEntity {
	var success []int
	var failed []int
	for _, objectId := range ids {
		conf := configuration.Configuration{}
		conf.ID = uint(objectId)
		response := conf.Delete()
		if response.Status != 200 {
			failed = append(failed, objectId)
		} else {
			success = append(success, objectId)
		}
	}
	if len(failed) == 0 {
		return http_utils.NewServiceResponse(http.StatusOK, gin.H{"Success": success})
	} else {
		return http_utils.NewServiceResponse(http.StatusPartialContent, gin.H{"Success": success, "failed": failed})
	}

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

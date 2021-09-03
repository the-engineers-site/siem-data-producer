package configuration

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/constants"
	"gitlab.com/yjagdale/siem-data-producer/database"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"net/http"
)

func (config *Configuration) Save() *http_utils.ResponseEntity {
	db, err := database.GetDBConnection()
	if err != nil {
		log.Errorln("Error while saving config.", err)
		return http_utils.NewInternalServerError("Unable to save to database", err)
	}
	db.Model(&Configuration{}).Save(config)
	return http_utils.NewOkResponse(constants.ResponseSave)
}

func (config *Configuration) Get() *http_utils.ResponseEntity {
	db, _ := database.GetDBConnection()
	err := db.Model(&Configuration{}).First(&config).Error
	if err != nil {
		log.Infoln("Error while fetching config")
		return http_utils.NewInternalServerError("Error while fetching config", err)
	}
	return http_utils.NewServiceResponse(http.StatusOK, config)
}

func (config *Configuration) GetAll() *http_utils.ResponseEntity {
	db, _ := database.GetDBConnection()
	var configurations []Configuration
	if database.ValidateConnection() {
		err := db.Model(&Configuration{}).Find(&configurations).Error
		if err != nil {
			log.Infoln("Error while fetching config")
			return http_utils.NewInternalServerError("Error while fetching config", err)
		}
		return http_utils.NewServiceResponse(http.StatusOK, configurations)
	}
	return http_utils.NewInternalServerError("DB Connection Error", nil)
}

func (config *Configuration) Update() *http_utils.ResponseEntity {
	return http_utils.NewOkResponse(constants.ResponseUpdate)
}

func (config *Configuration) Delete() *http_utils.ResponseEntity {
	db, err := database.GetDBConnection()
	if database.ValidateConnection() {
		err := db.Delete(config)
		if err != nil {
			log.Errorln("Error while deleting the config")
			return http_utils.NewServiceResponse(http.StatusBadRequest, gin.H{"Message": "Unable to delete record"})
		} else {
			return http_utils.NewOkResponse(constants.ResponseDelete)
		}
	} else {
		return http_utils.NewInternalServerError("Unable to connect to db", err)
	}
}

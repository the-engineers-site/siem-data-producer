package configuration

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/database"
	"net/http"
	"strings"
)

func (config *Configuration) Save() Response {
	var resp Response
	db, err := database.GetDBConnection()

	if err != nil {
		log.Errorln("Error while saving config.", err)
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	dbResponse := db.Model(&Configuration{}).Create(&config).Error
	if dbResponse != nil && strings.Contains(dbResponse.Error(), "UNIQUE constraint failed") {
		resp.SetMessage(http.StatusBadRequest, nil, gin.H{"reason": "Override key already exists", "code": 1002})
		return resp
	} else if dbResponse != nil {
		resp.SetMessage(http.StatusInternalServerError, nil, dbResponse.Error())
		return resp
	}
	resp.SetMessage(http.StatusCreated, config, nil)
	return resp
}

func (config *Configuration) Get() Response {
	db, err := database.GetDBConnection()
	var resp Response
	dbError := db.Model(&Configuration{}).First(&config).Error
	if dbError != nil {
		log.Infoln("Error while fetching config")
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	resp.SetMessage(http.StatusInternalServerError, nil, err)
	return resp
}

func (config *Configuration) GetAll() Response {
	var resp = Response{}
	db, err := database.GetDBConnection()
	var configurations []Configuration
	if database.ValidateConnection() {
		err := db.Model(&Configuration{}).Find(&configurations).Error
		if err != nil {
			log.Infoln("Error while fetching config")
			resp.SetMessage(http.StatusInternalServerError, nil, err)
			return resp
		}
		resp.SetMessage(http.StatusOK, configurations, nil)
		return resp
	}
	resp.SetMessage(http.StatusInternalServerError, "DB Connection Error", err)
	return resp
}

func (config *Configuration) Update() Response {
	var resp Response
	resp.SetMessage(http.StatusOK, "Updated successfully", nil)
	return resp
}

func (config *Configuration) Delete() Response {
	db, err := database.GetDBConnection()
	var resp Response
	if database.ValidateConnection() {
		err := db.Delete(config).RowsAffected
		if err == 0 {
			log.Errorln("Configuration not found. Affected rows ", err)
			resp.SetMessage(http.StatusNotFound, nil, errors.New("configuration not found"))
			return resp
		} else {
			resp.SetMessage(http.StatusOK, "Configuration deleted successfully", nil)
			return resp
		}
	} else {
		resp.SetMessage(http.StatusInternalServerError, "Internal server Error", err)
		return resp
	}
}

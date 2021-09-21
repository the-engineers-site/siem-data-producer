package producer

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/database"
	"net/http"
	"strings"
)

func (producerObject *Producer) Save() Response {
	var resp Response
	db, err := database.GetDBConnection()

	if err != nil {
		log.Errorln("Error while saving config.", err)
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	dbResponse := db.Model(&Producer{}).Create(&producerObject).Error
	if dbResponse != nil && strings.Contains(dbResponse.Error(), "UNIQUE constraint failed") {
		resp.SetMessage(http.StatusBadRequest, nil, gin.H{"reason": "Override key already exists", "code": 1002})
		return resp
	} else if dbResponse != nil {
		resp.SetMessage(http.StatusInternalServerError, nil, dbResponse.Error())
		return resp
	}
	resp.SetMessage(http.StatusCreated, producerObject, nil)
	return resp
}

func (producerObject *Producer) ExecutionsForProfile() ([]Producer, error) {
	var (
		producers []Producer
	)
	db, err := database.GetDBConnection()
	if err != nil {
		return nil, err
	}
	dbError := db.Model(&Producer{}).Where("profile_name=?", producerObject.ProfileName).Find(&producers).Error
	if dbError != nil {
		log.Infoln("Error while fetching config")
		return nil, dbError
	}
	return producers, nil
}

func (producerObject *Producer) Get() Response {
	db, err := database.GetDBConnection()
	var resp Response
	dbError := db.Model(&Producer{}).First(&producerObject).Error
	if dbError != nil {
		log.Infoln("Error while fetching config")
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	resp.SetMessage(http.StatusOK, producerObject, nil)
	return resp
}

func (producerObject *Producer) GetAll() Response {
	var resp = Response{}
	db, err := database.GetDBConnection()
	var producers []Producer
	if database.ValidateConnection() {
		err := db.Model(&Producer{}).Find(&producers).Error
		if err != nil {
			log.Infoln("Error while fetching config")
			resp.SetMessage(http.StatusInternalServerError, nil, err)
			return resp
		}
		resp.SetMessage(http.StatusOK, producers, nil)
		return resp
	}
	resp.SetMessage(http.StatusInternalServerError, "DB Connection Error", err)
	return resp
}

func (producerObject *Producer) Update() Response {
	var resp Response
	resp.SetMessage(http.StatusOK, "Updated successfully", nil)
	return resp
}

func (producerObject *Producer) Delete() Response {
	db, err := database.GetDBConnection()
	var resp Response
	if database.ValidateConnection() {
		err := db.Delete(producerObject).RowsAffected
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

package profile

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/database"
	"net/http"
	"strings"
)

func (profile *Profile) Save() Response {
	var resp Response
	db, err := database.GetDBConnection()

	if err != nil {
		log.Errorln("Error while saving config.", err)
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	dbResponse := db.Model(&Profile{}).Create(&profile).Error
	if dbResponse != nil && strings.Contains(dbResponse.Error(), "UNIQUE constraint failed") {
		resp.SetMessage(http.StatusBadRequest, nil, gin.H{"reason": "profile key already exists", "code": 1002})
		return resp
	} else if dbResponse != nil {
		resp.SetMessage(http.StatusInternalServerError, nil, dbResponse.Error())
		return resp
	}
	resp.SetMessage(http.StatusCreated, profile, nil)
	return resp
}

func (profile *Profile) Get() Response {
	db, err := database.GetDBConnection()
	var resp Response
	dbError := db.Model(&Profile{}).First(&profile).Error
	if dbError != nil {
		log.Infoln("Error while fetching config")
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	resp.SetMessage(http.StatusOK, profile, nil)
	return resp
}

func (profile *Profile) GetProfileByName() Response {
	db, err := database.GetDBConnection()
	var resp Response
	dbError := db.Model(&Profile{}).Where("name=?", profile.Name).First(&profile).Error
	if dbError != nil {
		log.Infoln("Error while fetching config")
		resp.SetMessage(http.StatusInternalServerError, nil, err)
		return resp
	}
	resp.SetMessage(http.StatusOK, profile, nil)
	return resp
}

func (profile *Profile) GetAll() Response {
	var resp = Response{}
	db, err := database.GetDBConnection()
	var configurations []Profile
	if database.ValidateConnection() {
		err := db.Model(&Profile{}).Find(&configurations).Error
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

func (profile *Profile) Update() Response {
	var resp Response
	resp.SetMessage(http.StatusOK, "Updated successfully", nil)
	return resp
}

func (profile *Profile) Delete() Response {
	db, err := database.GetDBConnection()
	var resp Response
	if database.ValidateConnection() {
		err := db.Delete(profile).RowsAffected
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

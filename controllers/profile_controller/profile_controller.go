package profile_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"siem-data-producer/constants"
	"siem-data-producer/models/configuration"
	"siem-data-producer/models/profile"
	"siem-data-producer/services"
	"siem-data-producer/utils/http_utils"
)

func SaveProfile(c *gin.Context) {
	var resp = profile.Response{}
	var profileObject profile.Profile

	err := c.ShouldBindJSON(&profileObject)
	if err != nil {
		resp.SetMessage(http.StatusBadRequest, "Invalid body", err.Error())
	} else {
		log.Infoln("Storing configuration")
		resp = services.ProfileService.SaveProfile(&profileObject)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func UpdateProfile(c *gin.Context) {
	var resp = configuration.Response{}
	var profileObject configuration.Configuration

	err := c.ShouldBindJSON(&profileObject)
	if err != nil {
		resp.SetMessage(http.StatusBadRequest, "Invalid Body", err)
	} else {
		log.Infoln("Storing configuration")
		resp = services.ConfigurationService.UpdateConfig(&profileObject)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func DeleteProfile(c *gin.Context) {
	var resp = &profile.Response{}
	var profileObject []string
	id := c.Param("id")

	if id != "" {
		log.Infoln("Deleting configuration with object", id)
		profileObject = append(profileObject, id)
	} else {
		err := c.ShouldBindJSON(&profileObject)
		if err != nil {
			c.JSON(http.StatusBadRequest, http_utils.NewOkResponse(constants.ResponseBadRequest+err.Error()))
			return
		}
		log.Infoln("Deleting object", profileObject)

	}
	resp = services.ProfileService.DeleteProfile(profileObject)
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func GetProfile(c *gin.Context) {
	var resp = profile.Response{}
	log.Infoln("Fetching configuration")
	id := c.Param("id")
	if id == "" {
		resp = services.ProfileService.GetProfile(&profile.Profile{})
	} else {
		profileObject := profile.Profile{}
		profileObject.Name = id
		resp = services.ProfileService.GetProfile(&profileObject)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

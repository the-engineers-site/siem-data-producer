package producer_controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"siem-data-producer/constants"
	"siem-data-producer/models/producer"
	"siem-data-producer/services"
	"siem-data-producer/utils/http_utils"
)

func StartProduce(c *gin.Context) {

	var resp = producer.Response{}
	var producerObject producer.Producer

	err := c.BindJSON(&producerObject)
	if err != nil {
		resp.SetMessage(http.StatusBadRequest, "Invalid body", err.Error())
	} else {
		log.Infoln("starting producer")
		resp = services.ProducerService.StartProducer(&producerObject)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return

}

func DeleteProfile(c *gin.Context) {
	var resp = producer.Response{}
	var producerObject []string
	id := c.Param("id")

	if id != "" {
		log.Infoln("Deleting configuration with object", id)
		producerObject = append(producerObject, id)
	} else {
		err := c.ShouldBindJSON(&producerObject)
		if err != nil {
			c.JSON(http.StatusBadRequest, http_utils.NewOkResponse(constants.ResponseBadRequest+err.Error()))
			return
		}
		log.Infoln("Deleting object", producerObject)

	}
	resp = services.ProducerService.DeleteProducer(producerObject)
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

func GetProduce(c *gin.Context) {
	var resp = producer.Response{}
	log.Infoln("Fetching producer")
	id := c.Param("id")
	if id == "" {
		resp = services.ProducerService.GetProducer(&producer.Producer{})
	} else {
		producerObject := producer.Producer{}
		producerObject.ExecutionId = id
		resp = services.ProducerService.GetProducer(&producerObject)
	}
	c.JSON(resp.GetStatus(), resp.GetResponse())
	return
}

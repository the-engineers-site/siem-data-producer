package services

import (
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/models/producer"
	"gitlab.com/yjagdale/siem-data-producer/models/profile"
	"gitlab.com/yjagdale/siem-data-producer/network_utils"
	"net/http"
)

var (
	ProducerService producerServiceInterface = &producerService{}
)

type producerServiceInterface interface {
	StartProducer(*producer.Producer) producer.Response
	StopProducer(*producer.Producer) producer.Response
	GetProducer(*producer.Producer) producer.Response
	DeleteProducer([]string) producer.Response
}

type producerService struct {
}

func (p producerService) DeleteProducer(ids []string) producer.Response {
	var statusCount struct {
		Success        []string      `json:"success,omitempty"`
		Failed         []string      `json:"failed,omitempty"`
		NotFount       []string      `json:"notFount,omitempty"`
		ProducerExists []interface{} `json:"producer_exists,omitempty"`
	}

	for _, id := range ids {
		var producerEntity = producer.Producer{
			ExecutionId: id,
		}
		response := producerEntity.Delete()
		if response.Status == http.StatusOK {
			statusCount.Success = append(statusCount.Success, id)
		} else if response.Status == http.StatusNotFound {
			statusCount.NotFount = append(statusCount.NotFount, id)
		} else if response.Status == http.StatusInternalServerError {
			statusCount.Failed = append(statusCount.Failed, id)
		}
	}
	resp := producer.Response{}
	resp.SetMessage(http.StatusOK, statusCount, nil)
	return resp
}

func (p producerService) GetProducer(producerObject *producer.Producer) producer.Response {
	if producerObject.ExecutionId == "" {
		return producerObject.GetAll()
	}
	return producerObject.Get()
}

func (p producerService) StartProducer(producerObject *producer.Producer) producer.Response {
	profileObj := profile.Profile{Name: producerObject.ProfileName}
	profileObj.Get()
	if profileObj.FilePath == "" {
		log.Info("No profile found. ")
		resp := producer.Response{}
		resp.SetMessage(http.StatusBadRequest, gin.H{"message": "Profile does not exists"}, nil)
		return resp
	}
	producerObject.Profile = &profileObj
	u, err := uuid.NewV4()
	if err == nil {
		producerObject.ExecutionId = u.String()
	}

	resp := producerObject.Save()
	if resp.Status != 201 {
		return resp
	} else {
		resp := network_utils.StartProducer(producerObject)
		return resp
	}
}

func (p producerService) StopProducer(producerObject *producer.Producer) producer.Response {
	log.Infoln(producerObject)
	panic("implement me")
}

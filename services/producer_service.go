package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"siem-data-producer/models/producer"
	"siem-data-producer/models/profile"
	"siem-data-producer/network_utils"
	"siem-data-producer/producectl/log_utils"
)

var (
	ProducerService producerServiceInterface = &producerService{}
)

type producerServiceInterface interface {
	StartProducer(*producer.Producer) producer.Response
	StopProducer(*producer.Producer) producer.Response
	GetProducer(*producer.Producer) producer.Response
	DeleteProducer([]string) producer.Response
	Init()
}

type producerService struct {
}

func (p producerService) Init() {
	var producerEntity = producer.Producer{}
	allProducer, err := producerEntity.GetAll()
	if err != nil {
		log.Errorln("Error while initiating executions")
	}

	for index, entity := range allProducer {
		log.Infoln("Init for ", index, " ID:", entity.ExecutionId)
		p.StartProducer(&entity)
	}
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
	response := producer.Response{}
	if producerObject.ExecutionId == "" {
		entities, err := producerObject.GetAll()
		if err != nil {
			response.SetMessage(http.StatusInternalServerError, nil, err)
			return response
		} else {
			response.SetMessage(http.StatusOK, entities, nil)
			return response
		}
	}
	return producerObject.Get()
}

func (p producerService) StartProducer(producerObject *producer.Producer) producer.Response {
	profileObj := profile.Profile{Name: producerObject.ProfileName}
	producerCtlCommand := "producerctl "
	profileObj.Get()
	if profileObj.FilePath == "" {
		log.Info("No profile found. ")
		resp := producer.Response{}
		resp.SetMessage(http.StatusBadRequest, gin.H{"message": "Profile does not exists"}, nil)
		return resp
	}

	producerObject.Profile = &profileObj

	if producerObject.Continues {
		producerCtlCommand = fmt.Sprintf(producerCtlCommand + "continues")
	} else {
		producerCtlCommand = producerCtlCommand + "once "
	}

	producerCtlCommand = fmt.Sprintf(producerCtlCommand+"--server=%s --protocol=%s --file_path=%s --eps=%d", profileObj.Destination, profileObj.Protocol, profileObj.FilePath, producerObject.Eps)

	log_utils.Log.Infoln("Starting process ", producerCtlCommand)

	if producerObject.ExecutionId != "" {
		log.Infoln("Restarted producer ", producerObject)
		resp := network_utils.StartProducer(producerObject)
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

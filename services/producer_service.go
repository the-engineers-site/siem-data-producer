package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"siem-data-producer/models/producer"
	"siem-data-producer/models/profile"
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
		resp := producerEntity.Get()
		if resp.Status == 200 {
			response := producerEntity.Delete()
			if response.Status == http.StatusOK {
				statusCount.Success = append(statusCount.Success, id)
			} else if response.Status == http.StatusNotFound {
				statusCount.NotFount = append(statusCount.NotFount, id)
			} else if response.Status == http.StatusInternalServerError {
				statusCount.Failed = append(statusCount.Failed, id)
			}
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
	producerCtlCommand := "producerctl"
	printCommand := "producerctl "
	var producerArgs []string
	profileObj.Get()
	if profileObj.FilePath == "" {
		log.Info("No profile found. ")
		resp := producer.Response{}
		resp.SetMessage(http.StatusBadRequest, gin.H{"message": "Profile does not exists"}, nil)
		return resp
	}

	producerObject.Profile = &profileObj

	if producerObject.Continues {
		producerArgs = append(producerArgs, "continues")
		printCommand = fmt.Sprintf(printCommand + "continues")
	} else {
		producerArgs = append(producerArgs, "once")
		producerCtlCommand = producerCtlCommand + "once "
	}

	producerArgs = append(producerArgs, "--server="+profileObj.Destination)
	producerArgs = append(producerArgs, "--protocol="+profileObj.Protocol)
	producerArgs = append(producerArgs, fmt.Sprintf("--file_path='%s'", profileObj.FilePath))
	producerArgs = append(producerArgs, fmt.Sprintf("--eps=%d", producerObject.Eps))
	printCommand = fmt.Sprintf("%s --server=%s --protocol=%s --file_path='%s' --eps=%d", printCommand, profileObj.Destination, profileObj.Protocol, profileObj.FilePath, producerObject.Eps)

	log_utils.Log.Infoln("Starting process ", printCommand)
	producerObject.Command = printCommand
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	process, err := os.StartProcess(producerCtlCommand, producerArgs, procAttr)
	fmt.Printf("%v\n", process)
	if err != nil {
		log_utils.Log.Errorf("ERROR Unable to run %s: %s\n", producerCtlCommand, err.Error())
	} else {
		log_utils.Log.Infof("%s running as pid %d\n", producerCtlCommand, process.Pid)
		producerObject.ProcessId = process.Pid
	}

	if producerObject.ExecutionId != "" {
		log.Infoln("Restarted producer ", producerObject)
		resp := producerObject.Update()
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
		resp := producerObject.Save()
		return resp
	}
}

func (p producerService) StopProducer(producerObject *producer.Producer) producer.Response {
	log.Infoln(producerObject)
	panic("implement me")
}

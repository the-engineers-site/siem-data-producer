package services

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/models/producer"
	"gitlab.com/yjagdale/siem-data-producer/models/profile"
	"net/http"
)

var (
	ProfileService ProfileServiceInterface = &profileService{}
)

type ProfileServiceInterface interface {
	SaveProfile(profile *profile.Profile) profile.Response
	UpdateProfile(profile *profile.Profile) profile.Response
	DeleteProfile(profile []string) *profile.Response
	GetProfile(profile *profile.Profile) profile.Response
}

type profileService struct {
}

func (c *profileService) UpdateProfile(profile *profile.Profile) profile.Response {
	return profile.Update()
}

func (c *profileService) DeleteProfile(ids []string) *profile.Response {
	var statusCount struct {
		Success        []string      `json:"success,omitempty"`
		Failed         []string      `json:"failed,omitempty"`
		NotFount       []string      `json:"notFount,omitempty"`
		ProducerExists []interface{} `json:"producer_exists,omitempty"`
	}
	var resp = profile.Response{}
	for _, objectId := range ids {
		p := profile.Profile{}
		p.Name = objectId
		var producerEntity producer.Producer
		producerEntity.ProfileName = objectId
		existenceResp, err := producerEntity.ExecutionsForProfile()
		if err != nil {
			statusCount.Failed = append(statusCount.Failed, objectId)
		} else {
			if len(existenceResp) != 0 {
				statusCount.ProducerExists = append(statusCount.ProducerExists, existenceResp)
			} else {
				response := p.Delete()
				if response.GetStatus() == 404 {
					statusCount.NotFount = append(statusCount.NotFount, objectId)
				} else if response.Status != 200 {
					statusCount.Failed = append(statusCount.Failed, objectId)
				} else {
					statusCount.Success = append(statusCount.Success, objectId)
				}
			}
		}
	}
	resp.SetMessage(http.StatusOK, statusCount, nil)
	return &resp
}

func (c *profileService) GetProfile(profile *profile.Profile) profile.Response {
	if profile.Name == "" {
		return profile.GetAll()
	}
	return profile.Get()
}

func (c profileService) SaveProfile(profileObject *profile.Profile) profile.Response {
	var resp profile.Response
	if profileObject.Validate() != nil {
		log.Debugln("Validation failed for configuration.")
		err := profileObject.Validate()
		resp.SetMessage(http.StatusBadRequest, "validation failed", err)

	}
	return profileObject.Save()
}

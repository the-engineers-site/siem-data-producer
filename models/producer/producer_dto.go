package producer

import (
	"errors"
	"siem-data-producer/models/profile"
)

type Producer struct {
	Profile     *profile.Profile `json:"profile,omitempty" gorm:"foreignKey:Name;references:ProfileName"`
	ExecutionId string           `json:"execution_id" gorm:"primarykey;not null"`
	ProfileName string           `json:"profile_name" binding:"required"`
	Eps         int              `json:"eps" binding:"required"`
	Continues   bool             `json:"continues" binding:"required"`
}

func (producerObject *Producer) Validate() error {
	if producerObject.Eps == 0 && producerObject.Continues {
		return errors.New("EPS is mandatory for continues producer")
	}

	if producerObject.ProfileName == "" {
		return errors.New("valid profile name needed")
	}
	return nil
}

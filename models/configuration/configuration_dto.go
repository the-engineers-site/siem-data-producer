package configuration

import (
	"github.com/lib/pq"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model
	OverrideKey    string         `json:"override_key" gorm:"not null"`
	OverrideValues pq.StringArray `json:"override_values" gorm:"type:text[]"`
}

func (config *Configuration) Validate() *http_utils.ResponseEntity {
	if config.OverrideKey == "" {
		return http_utils.NewBadRequestResponse("Override key is not provided")
	}

	if len(config.OverrideValues) <= 0 {
		return http_utils.NewBadRequestResponse("Override values is not provided")
	}

	return nil
}

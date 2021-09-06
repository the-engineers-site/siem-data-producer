package configuration

import (
	"errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model
	OverrideKey    string         `json:"override_key" gorm:"not null"`
	OverrideValues pq.StringArray `json:"override_values" gorm:"type:text[]"`
}

func (config *Configuration) Validate() error {
	if config.OverrideKey == "" {
		return errors.New("override key is not provided")
	}

	if len(config.OverrideValues) <= 0 {
		return errors.New("override values is not provided")
	}

	return nil
}

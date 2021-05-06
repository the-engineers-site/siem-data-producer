package configuration

import (
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"gorm.io/gorm"
)

type Configuration struct {
	gorm.Model
	LogFilePath string `json:"log_file_path"`
}

func (config *Configuration) Validate() *http_utils.ResponseEntity {
	if config.LogFilePath == "" {
		return http_utils.NewBadRequestResponse("File path cannot be empty")
	}
	return nil
}

package file_upload

import (
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"regexp"
)

type FileUpload struct {
	gorm.Model
	DeviceType   string
	DeviceVendor string
	File         *multipart.FileHeader `gorm:"type:blob"`
}

var validFormats = map[string]bool{
	".csv": true,
	".log": true,
}

func (fileUploadObject *FileUpload) Validate() *http_utils.ResponseEntity {
	if fileUploadObject.DeviceVendor == "" {
		return http_utils.NewBadRequestResponse("Device Vendor is empty")
	}

	if fileUploadObject.DeviceType == "" {
		return http_utils.NewBadRequestResponse("Device Type is empty")
	}

	if !validFormats[filepath.Ext(fileUploadObject.File.Filename)] {
		return http_utils.NewBadRequestResponse("Format is not supported")
	}
	reg := regexp.MustCompile("^/[[:print:]]+(/[[:print:]]+)*$")

	if reg.MatchString(filepath.Base(fileUploadObject.File.Filename)) {
		return http_utils.NewBadRequestResponse("File Name is not expected")
	}
	return nil
}

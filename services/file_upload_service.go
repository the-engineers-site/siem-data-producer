package services

import (
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"path/filepath"
	"siem-data-producer/models/file_upload"
	"siem-data-producer/utils/http_utils"
)

var FileUploadService fileUploadInterface = &fileUploadService{}

type fileUploadInterface interface {
	UploadFile(string, string, *multipart.FileHeader) *http_utils.ResponseEntity
}

type fileUploadService struct {
}

func (h *fileUploadService) UploadFile(deviceType string, deviceVendor string, file *multipart.FileHeader) *http_utils.ResponseEntity {
	log.Infoln("Processing request for upload file.", deviceType, "vendor:", deviceVendor)

	filename := filepath.Base(file.Filename)
	fileExtension := filepath.Ext(filename)

	log.Debugln("File Name:", filename, "and extension: ", fileExtension)

	uploader := file_upload.FileUpload{}
	uploader.File = file
	uploader.DeviceType = deviceType
	uploader.DeviceVendor = deviceVendor
	validateError := uploader.Validate()
	if validateError == nil {
		return uploader.Upload()
	} else {
		return validateError
	}

}

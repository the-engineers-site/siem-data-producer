package file_upload

import (
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"siem-data-producer/database"
	"siem-data-producer/utils/fileUtils"
	"siem-data-producer/utils/http_utils"
)

func (fileUploadObject *FileUpload) Upload() *http_utils.ResponseEntity {
	db, err := database.GetDBConnection()
	path, err := fileUtils.CreateOutputFolder(fileUploadObject.DeviceType, fileUploadObject.DeviceVendor)
	if err == nil {
		log.Infoln("Directory Created. Copying file contents")
		file, err := fileUploadObject.File.Open()
		if err == nil {
			defer func(file multipart.File) {
				err := file.Close()
				if err != nil {

				}
			}(file)
			outputFile, err := fileUtils.CreateFile(path + "/" + fileUploadObject.File.Filename)
			if err == nil {
				defer func(outputFile *os.File) {
					err := outputFile.Close()
					if err != nil {

					}
				}(outputFile)
				_, err = io.Copy(outputFile, file)
				if err != nil {
					return http_utils.NewInternalServerError("Failed to upload file", err)
				}

				err := db.Save(&UploadedFile{
					DeviceType:   fileUploadObject.DeviceType,
					DeviceVendor: fileUploadObject.DeviceVendor,
					Path:         path + "/" + fileUploadObject.File.Filename,
				})
				if err.Error != nil {
					log.Errorln("Error while storing to db", err.Error)
					resp := http_utils.NewInternalServerError("Unexpected Error", err.Error)
					fileUploadObject.rollbackChanges(path)
					return resp
				}
				return http_utils.NewOkResponse("File Uploaded success")
			}
		}
	}
	return http_utils.NewInternalServerError("Error while storing file content", err)
}

func (fileUploadObject *FileUpload) rollbackChanges(path string) {
	log.Infoln("Rolling back changes")

	err := fileUtils.RemoveFile(path + "/" + fileUploadObject.File.Filename)

	if err != nil {
		log.Errorln("Error while rolling back", err)
	}
}

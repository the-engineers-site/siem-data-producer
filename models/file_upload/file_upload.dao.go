package file_upload

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/utils/fileUtils"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
	"io"
	"mime/multipart"
	"os"
)

func (fileUploadObject *FileUpload) Upload() *http_utils.ResponseEntity {
	var err error
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
				if err == nil {
					return http_utils.NewOkResponse("File Uploaded successfully")
				}
			}
		}
	}
	return http_utils.NewInternalServerError("Error while storing file content", err)
}

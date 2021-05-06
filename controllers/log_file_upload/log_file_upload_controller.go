package log_file_upload

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/services"
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
)

func UploadFile(c *gin.Context) {

	var resp *http_utils.ResponseEntity

	deviceType := c.PostForm("Device Type")
	deviceVendor := c.PostForm("Device Vendor")

	file, err := c.FormFile("file")
	if err != nil {
		log.Errorln("Error while reading file", err)
		resp = http_utils.NewBadRequestResponse("Failed to reading file.")
	} else {
		resp = services.FileUploadService.UploadFile(deviceType, deviceVendor, file)
	}

	c.JSON(resp.Status, resp.Response)
	return

}

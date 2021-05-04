package app

import (
	"gitlab.com/yjagdale/siem-data-producer/controllers/health"
	"gitlab.com/yjagdale/siem-data-producer/controllers/log_file_upload"
)

func mapUrls() {
	router.GET("/ping", health.Ping)

	// file upload controllers
	router.Static("/ui", "./static")
	router.POST("/file/upload", log_file_upload.UploadFile)
}

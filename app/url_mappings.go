package app

import (
	"gitlab.com/yjagdale/siem-data-producer/controllers/configuration_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/health"
	"gitlab.com/yjagdale/siem-data-producer/controllers/log_file_upload"
)

func mapUrls() {
	router.GET("/ping", health.Ping)

	// file upload controllers
	router.Static("/ui", "./static")
	router.POST("/file/upload", log_file_upload.UploadFile)

	// configuration controllers

	configurationRoutes := router.Group("/configuration")
	configurationRoutes.GET("/", configuration_controller.GetConfiguration)
	configurationRoutes.GET("/:id", configuration_controller.GetConfiguration)
	configurationRoutes.POST("/", configuration_controller.SaveConfiguration)
	configurationRoutes.DELETE("/", configuration_controller.DeleteConfiguration)
	configurationRoutes.PUT("/", configuration_controller.UpdateConfiguration)
}

package app

import (
	"gitlab.com/yjagdale/siem-data-producer/controllers/configuration_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/health_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/log_file_upload"
	"gitlab.com/yjagdale/siem-data-producer/controllers/producer_controller"
)

func mapUrls() {
	router.GET("/ping", health_controller.Ping)

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

	// health controllers
	router.GET("/health", health_controller.Health)

	// producer controller
	producerController := router.Group("/produce")
	producerController.GET("/", producer_controller.Produce)

}

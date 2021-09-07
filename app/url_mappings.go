package app

import (
	"github.com/thinkerou/favicon"
	"gitlab.com/yjagdale/siem-data-producer/controllers/configuration_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/health_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/log_file_upload"
	"gitlab.com/yjagdale/siem-data-producer/controllers/producer_controller"
)

func mapUrls() {
	router.GET("/ping", health_controller.Ping)

	// file upload controllers
	router.Static("/ui", "./static")
	router.Use(favicon.New("./static/favicon.ico"))
	router.POST("/file/upload", log_file_upload.UploadFile)

	configurationMapping()
	health()
	producer()
}

func configurationMapping() {
	// configuration controllers
	configurationRoutes := router.Group("/configuration")
	configurationRoutes.GET("/", configuration_controller.GetConfiguration)
	configurationRoutes.GET("/:id", configuration_controller.GetConfiguration)
	configurationRoutes.POST("/", configuration_controller.SaveConfiguration)
	configurationRoutes.DELETE("/", configuration_controller.DeleteConfiguration)
	configurationRoutes.DELETE("/:id", configuration_controller.DeleteConfiguration)
	configurationRoutes.PUT("/", configuration_controller.UpdateConfiguration)
	configurationRoutes.POST("/reload", configuration_controller.ReloadConfiguration)
	configurationRoutes.GET("/overrides", configuration_controller.GetOverrides)
}

func health() {
	// health controllers
	router.GET("/health", health_controller.Health)
}

func producer() {
	// producer controller
	producerController := router.Group("/produce")
	producerController.GET("/", producer_controller.Produce)
}

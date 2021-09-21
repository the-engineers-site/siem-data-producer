package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/thinkerou/favicon"
	"gitlab.com/yjagdale/siem-data-producer/controllers/configuration_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/health_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/log_file_upload"
	"gitlab.com/yjagdale/siem-data-producer/controllers/producer_controller"
	"gitlab.com/yjagdale/siem-data-producer/controllers/profile_controller"
	_ "gitlab.com/yjagdale/siem-data-producer/docs"
	"gitlab.com/yjagdale/siem-data-producer/utils/utils"
)

var api *gin.RouterGroup

func mapUrls() {
	router.GET("/ping", health_controller.Ping)

	// file upload controllers
	router.Static("/ui", "./static")
	router.Use(favicon.New("./static/favicon.ico"))
	router.POST("/file/upload", log_file_upload.UploadFile)
	api = router.Group("/v1")
	configurationMapping()
	profileMapping()
	health()
	producerMapping()
	url := ginSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", utils.GetOutboundIP(), utils.GetPort()))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func configurationMapping() {
	// configuration controllers
	configurationRoutes := api.Group("/configuration")
	configurationRoutes.GET("/", configuration_controller.GetConfiguration)
	configurationRoutes.GET("/:id", configuration_controller.GetConfiguration)
	configurationRoutes.POST("/", configuration_controller.SaveConfiguration)
	configurationRoutes.DELETE("/", configuration_controller.DeleteConfiguration)
	configurationRoutes.DELETE("/:id", configuration_controller.DeleteConfiguration)
	configurationRoutes.PUT("/", configuration_controller.UpdateConfiguration)
	configurationRoutes.POST("/reload", configuration_controller.ReloadConfiguration)
	configurationRoutes.GET("/overrides", configuration_controller.GetOverrides)
}

func profileMapping() {
	// configuration controllers
	configurationRoutes := api.Group("/profile")
	configurationRoutes.GET("/", profile_controller.GetProfile)
	configurationRoutes.GET("/:id", profile_controller.GetProfile)
	configurationRoutes.POST("/", profile_controller.SaveProfile)
	configurationRoutes.DELETE("/", profile_controller.DeleteProfile)
	configurationRoutes.DELETE("/:id", profile_controller.DeleteProfile)
	configurationRoutes.PUT("/", profile_controller.UpdateProfile)
}

func health() {
	// health controllers
	router.GET("/health", health_controller.Health)
}

func producerMapping() {
	// producer controller
	producerController := api.Group("/produce")
	producerController.POST("/", producer_controller.StartProduce)
	producerController.GET("/", producer_controller.GetProduce)
	producerController.GET("/:id", producer_controller.GetProduce)
	producerController.DELETE("/", producer_controller.DeleteProfile)
	producerController.DELETE("/:id", producer_controller.DeleteProfile)

}

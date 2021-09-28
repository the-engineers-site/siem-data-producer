package app

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/thinkerou/favicon"
	"os"
	"siem-data-producer/controllers/configuration_controller"
	"siem-data-producer/controllers/health_controller"
	"siem-data-producer/controllers/log_file_upload"
	"siem-data-producer/controllers/producer_controller"
	"siem-data-producer/controllers/profile_controller"
	_ "siem-data-producer/docs"
)

var api *gin.RouterGroup

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func mapUrls() {
	router.Use(CORSMiddleware())
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
	var path string
	if os.Getenv("IP_ADDR") != "" {
		path = os.Getenv("IP_ADDR") + "/swagger/doc.json"
	} else {
		path = "http://localhost:8082/swagger/doc.json"
	}
	url := ginSwagger.URL(path)
	router.Use(CORSMiddleware()).GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
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

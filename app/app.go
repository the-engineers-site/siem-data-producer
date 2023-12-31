package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"siem-data-producer/services"
	"siem-data-producer/utils/utils"
	"time"
)

var (
	router = gin.New()
)

func init() {

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	log.SetReportCaller(true)
	initDBMigration()
	initReload()
	router.Use(gin.Recovery())
}

func initReload() {
	ticker := time.NewTicker(55 * time.Second)
	services.ConfigurationService.Reload()
	go func() {
		for {
			select {
			case <-ticker.C:
				services.ConfigurationService.Reload()
			}
		}
	}()
}

func initProducers() {
	services.ProducerService.Init()
}

func StartApplication() {
	log.Infoln("Start Application version x")
	mapUrls()
	fmt.Println("Execution started")
	log.Info("about to start the application...")

	log.Infoln("Starting application on :", utils.GetPort())
	initProducers()
	err := router.Run(":" + utils.GetPort())
	if err != nil {
		log.Fatalln("Error while starting service.", err)
	}
}

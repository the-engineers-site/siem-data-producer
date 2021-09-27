package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/services"
	"gitlab.com/yjagdale/siem-data-producer/utils/utils"
	"time"
)

var (
	router = gin.New()
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func init() {
	router.Use(CORSMiddleware())
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

func StartApplication() {
	mapUrls()
	fmt.Println("Execution started")
	log.Info("about to start the application...")

	log.Infoln("Starting application on :", utils.GetPort())
	err := router.Run(":" + utils.GetPort())
	if err != nil {
		log.Fatalln("Error while starting service.", err)
	}
}

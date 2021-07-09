package app

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/database"
	"gitlab.com/yjagdale/siem-data-producer/models/configuration"
	"gitlab.com/yjagdale/siem-data-producer/models/file_upload"
	"gitlab.com/yjagdale/siem-data-producer/models/health_models"
	"time"
)

func initDBMigration() {
	log.Infoln("Migration of database started", time.Now())
	db, err := database.GetDBConnection()

	if err != nil {
		log.Errorln("Error while connecting", err)
		panic(1)
	}

	migrationError := db.AutoMigrate(&configuration.Configuration{}, &file_upload.UploadedFile{}, &health_models.Health{})

	if migrationError != nil {
		log.Fatalln("Error while migrating database", migrationError)
	}
	log.Infoln("DB migration completed")
}

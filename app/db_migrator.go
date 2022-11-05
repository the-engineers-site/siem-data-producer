package app

import (
	log "github.com/sirupsen/logrus"
	"siem-data-producer/database"
	"siem-data-producer/models/configuration"
	"siem-data-producer/models/file_upload"
	"siem-data-producer/models/health_models"
	"siem-data-producer/models/producer"
	"siem-data-producer/models/profile"
	"time"
)

func initDBMigration() {
	log.Infoln("Migration of database started", time.Now())
	db, err := database.GetDBConnection()

	if err != nil {
		log.Errorln("Error while connecting", err)
		panic(1)
	}

	migrationError := db.AutoMigrate(
		&configuration.Configuration{},
		&profile.Profile{},
		&file_upload.UploadedFile{},
		&health_models.Health{},
		&producer.Producer{})

	if migrationError != nil {
		log.Fatalln("Error while migrating database", migrationError)
	}
	log.Infoln("DB migration completed")
}

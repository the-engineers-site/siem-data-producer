package database

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/yjagdale/siem-data-producer/models/health_models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var databaseConnection *gorm.DB

func GetDBConnection() (*gorm.DB, error) {
	return connectDB()
}

func connectDB() (*gorm.DB, error) {
	var err error
	log.Infoln("Connecting to database")
	databaseConnection, err = gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	log.Debugln("Connected ", err == nil)
	return databaseConnection, err
}

func ValidateConnection() bool {
	if databaseConnection == nil {
		_, err := connectDB()
		if err != nil {
			log.Errorln("Error while connecting to database", err)
			return false
		}
	}
	health := health_models.Health{}
	err := databaseConnection.Select(&health).Error
	if err != nil {
		return false
	}
	return true
}

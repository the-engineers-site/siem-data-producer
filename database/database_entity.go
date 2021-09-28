package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"siem-data-producer/constants"
	"siem-data-producer/models/health_models"
)

var databaseConnection *gorm.DB

func GetDBConnection() (*gorm.DB, error) {
	return connectDB()
}

func connectDB() (*gorm.DB, error) {
	var err error
	var dbPath string
	log.Debugln("Connecting to database")
	dbPath = os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = constants.DefaultDbPath + "/" + constants.DefaultDbName
		err = os.MkdirAll(constants.DefaultDbPath, 0777)
		if err != nil {
			log.Errorln("Error while creating database directory", err)
		}
	} else {
		log.Infoln("DB path provided in env, Using", dbPath)
		dbPath = dbPath + "/" + constants.DefaultDbName
	}

	databaseConnection, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	log.Debugln("Connected ", err == nil)
	return databaseConnection, err
}

func ValidateConnection() bool {
	if databaseConnection == nil {
		_, err := connectDB()
		if err != nil {
			log.Errorln("Error while connecting to database", err)
			return false
		} else {
			log.Debugln("Connection created successfully")
		}
	}
	health := health_models.Health{}
	err := databaseConnection.Model(health_models.Health{}).First(&health).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Errorln("Error while checking health", err)
		return false
	}
	log.Debugln("DB health validated successfully")
	return true
}

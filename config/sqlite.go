package config

import (
	"os"

	"github.com.br/thenopholo/uid_eveo/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating one....")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite Openig Error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.CompanyModel{})
	if err != nil {
		logger.Errorf("SQLite Automigration Error: %v", err)
		return nil, err
	}
	return nil, err
}
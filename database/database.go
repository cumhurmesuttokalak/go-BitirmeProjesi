package database

import (
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/helpers"
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func DatabaseConnection() {
	db, err := gorm.Open(sqlite.Open("bitirmeprojesidb"), &gorm.Config{})
	helpers.CheckErr(err)
	db.AutoMigrate(&models.User{})
	Database = DbInstance{Db: db}
}

package config

import (
	"Interfaceginrestapi/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {

	var urlDSN = "root:Quest1234@tcp(localhost:3306)/newdb?parseTime=true"
	var err error

	db, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&models.Company{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

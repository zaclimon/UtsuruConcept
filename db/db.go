package db

import (
	"UtsuruConcept/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var dbObj *gorm.DB

// Init initializes the database object.
func Init() {
	if dbObj == nil {
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?&parseTime=True&loc=UTC", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := gorm.Open("mysql", connectionString)

		if err != nil {
			fmt.Println("Error while trying to initialize the database", err)
		}

		dbObj = db
		dbObj.AutoMigrate(&models.User{}, &models.Image{}, &models.ImageData{})
	} else {
		fmt.Println("Database has already been initialized")
	}
}

// GetDB retrieves the database object after it has been initialized.
func GetDb() *gorm.DB {
	if dbObj != nil {
		return dbObj
	}
	fmt.Println("Please initialize database object")
	return nil
}

// Stop closes the connection with the database.
func Stop() {
	if dbObj != nil {
		defer dbObj.Close()
	}
}

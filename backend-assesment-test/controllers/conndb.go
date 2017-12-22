package controllers

import (
	"fmt"
	"os"

	"github.com/ervinismu/backend-assesment-test/models"
	"github.com/jinzhu/gorm"
	// for connect postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Function for connected db
func InitDb() *gorm.DB {

	var err error
	var user models.User
	var notes models.Notes

	host := os.Getenv("")
	if host == "" {
		host = "your-connection-db"
	}

	db, err := gorm.Open("postgres", host)
	db.AutoMigrate(user, notes)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Db connected!")
	}
	return db
}

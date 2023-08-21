package helper

import (
	"fmt"
	"log"

	"github.com/nehaal10/CardGaurd/models"
	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) {
	//checks if there already a table
	isThereFraud := db.Migrator().HasTable(&models.FraudDataBase{})
	isThereUser := db.Migrator().HasTable(&models.UserDatabase{})
	//if not create a table
	if !isThereFraud && !isThereUser {
		db.Migrator().CreateTable(&models.FraudDataBase{}, &models.UserDatabase{})
	}

	// this all will moved to another file
	// we will get the data from spi calls
	user := models.UserDatabase{UserName: "Nehaal", EmailID: "nehaal-gmail.com"}
	msg := user.ValidateEmail()
	if !msg {
		log.Fatal(msg)
		fmt.Println("In Valid Email ID")
	} else {
		db.Create(&user)
	}
}

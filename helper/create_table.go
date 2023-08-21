package helper

import (
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
}

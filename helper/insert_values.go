package helper

import (
	"fmt"
	"log"

	"github.com/nehaal10/CardGaurd/models"
	"gorm.io/gorm"
)

func Insert_users(user models.UserDatabase, db *gorm.DB) string {
	msg := user.ValidateEmail()
	if !msg {
		log.Fatal(msg)
		return "In Valid Email ID"
	} else {
		db.Create(&user)
	}
	str := fmt.Sprintf("%d inserted in the databse", user.ID)
	return str
}

func Insert_fraud(db *gorm.DB, fraud models.FraudDataBase) *gorm.DB {
	result := db.Create(&fraud)
	//str := fmt.Sprintf("%d inserted in the databse", fraud.ID)
	return result
}

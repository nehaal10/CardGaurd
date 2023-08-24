package helper

import (
	"fmt"

	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
	"gorm.io/gorm"
)

func Insert_users(user models.UserDatabase, db *gorm.DB) string {
	//hash password
	user.Password = string(utils.HashPassword(user.Password))
	result := db.Create(&user)
	isErr := utils.CheckErrReturn(result.Error)
	if !isErr {
		return fmt.Sprintf("Cannot Register--Try AGAIN--%s", result.Error)
	}
	return "Registerd"
}

func Insert_fraud(db *gorm.DB, fraud models.FraudDataBase) string {

	isCheck := Cvv(fraud.Cvv) && DateVerify(fraud.ExpiryDate) && Number(fraud.CreditCardNo)

	if !isCheck {
		result := db.Create(&fraud)
		utils.CheckError(result.Error)
		return "Not Verified"
	}
	return "Your are verified"
}

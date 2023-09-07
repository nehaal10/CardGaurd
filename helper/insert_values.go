package helper

import (
	"fmt"

	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
	"gorm.io/gorm"
)

func Insert_users(user models.UserDatabase, db *gorm.DB) (string, int) {
	//hash password
	user.Password = string(utils.HashPassword(user.Password))
	result := db.Create(&user)
	isErr := utils.CheckErrReturn(result.Error)
	if !isErr {
		return fmt.Sprintf("Cannot Register--Try AGAIN--%s", result.Error), 400
	}
	return "Registerd", 200
}

func Insert_fraud(db *gorm.DB, fraud models.FraudDataBase) string {
	//var wg sync.WaitGroup
	//var result = make(chan bool)
	//var isCheck = true
	isCheck := Cvv(fraud.Cvv) && DateVerify(fraud.ExpiryDate) && Number(fraud.CreditCardNo)
	/*wg.Add(3)

	go Cvv(fraud.Cvv, result)
	go DateVerify(fraud.ExpiryDate, result)
	go Number(fraud.CreditCardNo, result)

	wg.Wait()

	if !<-result {
		isCheck = false
	}*/

	if !isCheck {
		result := db.Create(&fraud)
		utils.CheckError(result.Error)
		return "Not Verified"
	}
	return "Your are verified"
}

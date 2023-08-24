package helper

import (
	"fmt"

	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
	"gorm.io/gorm"
)

func GetOne(db *gorm.DB, user models.UserDatabase) (string, int) {
	var u models.UserDatabase
	result := db.Find(&u, "user_name = ?", user.UserName)
	//checking for worong user
	if u.ID == 0 {
		return "User doesnt exist", 400
	}
	//any database issue
	isErr := utils.CheckErrReturn(result.Error)
	if !isErr {
		str := fmt.Sprintf("%s", result.Error)
		return str, 404
	}
	//check if they have same user name
	isSameUser := u.UserName == user.UserName
	if !isSameUser {
		return "invalid User", 400
	}
	isSame := utils.CompareHash(user.Password, []byte(u.Password))
	if !isSame {
		return "Invalid User", 400
	}

	return "Valid User", 200
}

package models

import (
	"net/mail"

	"github.com/nehaal10/CardGaurd/utils"
)

type FraudDataBase struct {
	ID         int64 `gorm:"primaryKey"`
	CreditCard string
}

type UserDatabase struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
	EmailID  string
}

func (u *UserDatabase) ValidateEmail() bool {
	_, err := mail.ParseAddress(u.EmailID)
	isVal := utils.CheckErrReturn(err)
	return isVal
}

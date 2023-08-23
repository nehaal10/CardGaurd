package models

import (
	"net/mail"

	"github.com/go-playground/validator"
	"github.com/nehaal10/CardGaurd/utils"
)

type FraudDataBase struct {
	ID           uint   `json:"id" validate:"numeric" gorm:"primaryKey"`
	CreditCardNo string `json:"creditcardno" validate:"required"`
	Cvv          string `json:"cvv" validate:"required"`
	ExpiryDate   string `json:"expirydate" validate:"required"`
}

func (f *FraudDataBase) Validating() bool {
	var val *validator.Validate = validator.New()
	err := val.Struct(f)
	isCheck := utils.CheckErrReturn(err)
	return isCheck
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

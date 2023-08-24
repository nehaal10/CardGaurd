package models

import (
	"net/mail"

	"github.com/go-playground/validator"
	"github.com/nehaal10/CardGaurd/utils"
)

type FraudDataBase struct {
	ID           uint   `json:"id" validate:"numeric" gorm:"primaryKey"`
	CreditCardNo string `json:"creditcardno" validate:"required" gorm:"unique"`
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
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserName string `json:"username" validate:"required" gorm:"unique"`
	EmailID  string `json:"emailid" validate:"required" gorm:"unique"`
	Password string `json:"password" validate:"required"`
}

func (u *UserDatabase) ValidateEmail() bool {
	_, err := mail.ParseAddress(u.EmailID)
	isVal := utils.CheckErrReturn(err)
	return isVal
}

func (f *UserDatabase) ValidatingUser() bool {
	var val *validator.Validate = validator.New()
	err := val.Struct(f)
	isCheck := utils.CheckErrReturn(err)
	return isCheck
}

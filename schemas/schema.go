package schemas

import (
	"github.com/go-playground/validator/v10"
	"github.com/nehaal10/CardGaurd/utils"
)

type UserInput struct {
	//ID uint `json:"id" validate:"numeric" gorm:"primaryKey"`
	CreditCardNo string `json:"creditcardno" validate:"required"`
	Cvv          string `json:"cvv" validate:"required"`
	ExpiryDate   string `json:"expirydate" validate:"required"`
}

func (f *UserInput) Validating() bool {
	var val *validator.Validate = validator.New()
	err := val.Struct(f)
	isCheck := utils.CheckErrReturn(err)
	return isCheck
}

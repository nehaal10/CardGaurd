package models

type FraudDataBase struct {
	ID         int64 `gorm:"primaryKey"`
	CreditCard string
}

type UserDatabase struct {
	ID       uint `gorm:"primaryKey"`
	UserName string
}

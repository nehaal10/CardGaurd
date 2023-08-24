package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/helper"
	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
	"gorm.io/gorm"
)

var db *gorm.DB = helper.DB

func Fraud_users(c *gin.Context) {
	var creditcardDetails models.FraudDataBase
	err := c.ShouldBindJSON(&creditcardDetails)
	check := creditcardDetails.Validating()

	if !check {
		c.Status(400)
		return
	}

	utils.CheckError(err)
	str := helper.Insert_fraud(db, creditcardDetails)
	c.JSON(200, str)
}

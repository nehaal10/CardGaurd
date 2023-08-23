package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/helper"
	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
)

func Fraud_users(c *gin.Context) {
	var creditcardDetails models.FraudDataBase
	err := c.ShouldBindJSON(&creditcardDetails)
	utils.CheckError(err)
	db := helper.DB
	fmt.Println(db)
	res := helper.Insert_fraud(db, creditcardDetails)
	if res.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, "Success")
}

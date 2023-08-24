package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/helper"
	"github.com/nehaal10/CardGaurd/models"
)

func Login(c *gin.Context) {
	var user models.UserDatabase
	c.ShouldBindJSON(&user)

	isValid := user.ValidateEmail() && user.ValidatingUser()
	if !isValid {
		c.JSON(200, "Invalid User")
	}
	str, num := helper.GetOne(db, user)

	//jwt

	c.JSON(num, str)
}

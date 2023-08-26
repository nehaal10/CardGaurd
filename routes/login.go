package routes

import (
	"time"

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

	if num != 200 {
		c.JSON(num, str)
	} else {
		//jwt
		val := helper.JwtAuth(user)

		c.SetCookie("user", val, int(time.Now().Add(time.Minute*5).Unix()), "/", "localhost", true, true)
		c.JSON(200, "Loged in")
	}
}

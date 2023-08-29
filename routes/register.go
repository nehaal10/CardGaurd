package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/helper"
	"github.com/nehaal10/CardGaurd/models"
)

func Register(c *gin.Context) {
	var userRegister models.UserDatabase
	c.ShouldBindJSON(&userRegister)
	check := userRegister.ValidateEmail() && userRegister.ValidatingUser()
	if !check {
		c.JSON(400, "connot register")
		return
	}

	str := helper.Insert_users(userRegister, db)
	c.JSON(200, str)
}

package routes

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {
	c.SetCookie("user", "", -1, "/", "localhost", true, true)
	c.JSON(200, "Logged out")
}

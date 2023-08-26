package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/routes"
)

func main() {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.POST("/post", routes.Fraud_users)
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)
	r.GET("/logout", routes.Logout)
	r.Run()
}

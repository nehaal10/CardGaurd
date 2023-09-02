package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/routes"
)

func main() {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.RedirectTrailingSlash = true
	//r.Use(static.Serve("/", static.LocalFile("/assets/build", true)))
	r.POST("/api/post", routes.Fraud_users)
	r.POST("/api/register", routes.Register)
	r.POST("/api/login", routes.Login)
	r.GET("/api/logout", routes.Logout)
	r.Run()
}

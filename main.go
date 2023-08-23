package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/CardGaurd/routes"
)

func main() {
	r := gin.Default()
	r.POST("/post", routes.Fraud_users)
	r.Run()
}

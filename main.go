package main

import (
	"github.com/nehaal10/CardGaurd/controller"
	"github.com/nehaal10/CardGaurd/helper"
)

func main() {
	db := controller.Init()
	helper.CreateTable(db)
}

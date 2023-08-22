package main

import (
	"fmt"

	"github.com/nehaal10/CardGaurd/helper"
)

func main() {
	//db := controller.Init()
	//helper.CreateTable(db)
	check := helper.Cvv("685")
	fmt.Println(check)
	isValid := helper.DateVerify("09/29")
	fmt.Println(isValid)
	isValidCredicardNum := helper.Number("4000-0566-5566-5556")
	fmt.Println(isValidCredicardNum)

}

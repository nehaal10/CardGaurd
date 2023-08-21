package utils

import (
	"fmt"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func CheckErrReturn(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

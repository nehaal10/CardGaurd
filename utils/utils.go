package utils

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func CheckErrReturn(err error) bool {
	return err == nil
}

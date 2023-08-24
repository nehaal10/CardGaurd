package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	CheckError(err)
	return hash
}

package controller

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nehaal10/CardGaurd/models"
	"github.com/nehaal10/CardGaurd/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	err := godotenv.Load("/Users/nehaalreddy/Developer/Projects/CardGaurd/.env")
	utils.CheckError(err)
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPassword, dbName)

	//connecting to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckError(err)
	fmt.Println("Successfully connected the database")

	//creating table
	isThereFraud := db.Migrator().HasTable(&models.FraudDataBase{})
	isThereUser := db.Migrator().HasTable(&models.UserDatabase{})
	//if not create a table
	if !isThereFraud && !isThereUser {
		db.Migrator().CreateTable(&models.FraudDataBase{}, &models.UserDatabase{})
	}

	return db
}

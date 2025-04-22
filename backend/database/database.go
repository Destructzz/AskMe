package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	err := godotenv.Load()

	if err != nil{
		log.Fatalf("Error to load env: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatalf("Error to connect database: %v", err)
	}

	log.Println("Database Connected")

}

func Migrate(){
	DB.AutoMigrate(&User{}, &Poll{}, &Card{}, &CardField{})
	log.Println("Migration complited")
}
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshua468/music-stream-api/config"
	"github.com/joshua468/music-stream-api/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvVariables() {
err:= godotenv.Load()
if err!= nil {
	log.Fatal("error connecting to .env file")
}	

}

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=% sslmode=disable",
	 os.Getenv("DB_HOST"),
	 os.Getenv("DB_USER"),
	 os.Getenv("DB_PASSWORD"),
	 os.Getenv("DB_NAME"),
	 os.Getenv("DB_PORT"))


	var err error
	database.DB,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!= nil {
		log.Fatal("failed to connect to database:",err)
	}
	fmt.Println("Database connection esthablished")
}

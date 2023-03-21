package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pg *gorm.DB
var Connected bool = false

var dsn string = getDbUrl()

func ConnectToDatabase() {
	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Pg = pg
	Connected = true

}
func GetDBConnection() *gorm.DB {
	if !Connected {
		ConnectToDatabase()
	}
	return Pg
}

func getDbUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbName)
}

package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new Datbase connection")

	//dbUsername := os.Getenv("DB_USERNAME")
	dbUsername := "postgres"
	//dbPassword := os.Getenv("DB_PASSWORD")
	dbPassword := "postgres"
	//dbHost := os.Getenv("DB_HOST")
	dbHost := "localhost"
	//dbTable := os.Getenv("DB_TABLE")
	dbTable := "postgres"
	//dbPort := os.Getenv("DB_PORT")
	dbPort := "5432"
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}
	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil

}

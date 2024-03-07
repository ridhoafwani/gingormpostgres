package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	DBName   = "postgres"
)

func DatabaseConnection() *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, DBName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	// helper.ErrorPanic(err)
	if err != nil {
		panic(err)
	}
	fmt.Println("? Connected Successfully to the Database")
	return db
}

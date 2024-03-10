package main

import (
	"fmt"

	"github.com/ridhoafwani/gingormpostgres/config"
	"github.com/ridhoafwani/gingormpostgres/models"
)

func main() {
	dbName := "orders_by"
	defaultDb := "postgres"
	db := config.DatabaseConnection(&defaultDb)

	var exists bool
	db.Raw("SELECT EXISTS(SELECT FROM pg_database WHERE datname = ?)", dbName).Scan(&exists)

	if !exists {
		db.Exec("CREATE DATABASE " + dbName)
	}

	db = config.DatabaseConnection(&dbName)

	db.AutoMigrate(&models.Item{}, &models.Orders{})
	fmt.Println("Migration complete")

}

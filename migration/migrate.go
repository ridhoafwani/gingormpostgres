package main

import (
	"fmt"
	"log"

	"github.com/ridhoafwani/gingormpostgres/config"
	"github.com/ridhoafwani/gingormpostgres/models"
)

func main() {
	db := config.DatabaseConnection()

	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.DBName)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		log.Fatal(rs.Error) 
	}

	// if not create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", "orders_by")
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Fatal(rs.Error)
		}

		// close db connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			log.Fatal(err)
		}
	}


	db.AutoMigrate(&models.Items{}, &models.Orders{})
	fmt.Println("Migration complete")
}
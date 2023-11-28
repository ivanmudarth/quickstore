package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBInit() {
	// Connect to MySQL instance
	var err error
	DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test_db")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Println("MySQL database initialized")
	}

	CreateAllTables()
}

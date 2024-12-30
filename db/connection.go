package db

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// connStr := "postgresql://root@localhost:26257/testing102?sslmode=disable"
	connStr := "postgresql://root@localhost:26257/techwave_writers?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	fmt.Println("Successfully connected to the database!")
}
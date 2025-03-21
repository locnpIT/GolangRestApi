package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/events?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Test the database connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Set database connection pool settings
	DB.SetConnMaxLifetime(3 * time.Minute)
	DB.SetMaxOpenConns(10) // how many connection can be open simultaneously
	DB.SetMaxIdleConns(10) //

	log.Println("Database connection established successfully")

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			location VARCHAR(255) NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INT
		);

	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}
}

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

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL);
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			location VARCHAR(255) NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INT,
			FOREIGN KEY (user_id) REFERENCES users(id)
        	ON DELETE CASCADE
		);

	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXIST registrations (
			id INT AUTO_INCREMENT PRIMARY KEY,
			event_id INT,
			user_id INT,
			FOREIGN KEY (user_id) REFERENCES users(id)
			FIREIGN KEY (event_id) REFERENCES events(id)
		)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create registration table")
	}

}

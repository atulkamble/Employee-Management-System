package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	// Load DB config from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Optional debug log
	fmt.Printf("📦 Connecting to PostgreSQL: host=%s port=%s dbname=%s user=%s\n", host, port, dbname, user)

	// Construct PostgreSQL connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Database ping error: %v", err)
	}

	fmt.Println("✅ Connected to the database successfully")
}

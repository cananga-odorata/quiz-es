package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to database successfully")

	// Read and execute migration file
	migrationSQL, err := os.ReadFile("migrations/000001_create_quizzes_table.up.sql")
	if err != nil {
		log.Fatalf("Failed to read migration file: %v", err)
	}

	_, err = db.Exec(string(migrationSQL))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v", err)
	}

	fmt.Println("Migration completed successfully!")
}

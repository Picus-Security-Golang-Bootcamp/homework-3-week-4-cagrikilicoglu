package main

import (
	"bookApp/pkg/Base/book"
	postgres "bookApp/pkg/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// Set environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init")
	}
	log.Println("Postgress connected")

	// Repositories
	bookrepo := book.NewBookRepository(db)
	bookrepo.Migrations()
	// bookList.CreateList()
}

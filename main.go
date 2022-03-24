package main

import (
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

	_, err = postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init")
	}
	log.Println("Postgress connected")

	// bookList.CreateList()
}
